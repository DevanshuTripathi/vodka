package mixers

import (
	"errors"
	"sync"
	"time"

	"github.com/DevanshuTripathi/vodka"
)

type limiter struct {
	lastUpdate time.Time
	tokens     float64
	rate       float64 // tokens per second
	burst      int     // max tokens
	mu         sync.Mutex
}

type visitor struct {
	l        *limiter
	lastSeen time.Time
}

type VodkaRateLimiter struct {
	visitors map[string]*visitor
	mu       sync.Mutex
	rate     float64
	burst    int

	done chan struct{}
	once sync.Once
}

func newLimiter(r float64, b int) *limiter {
	return &limiter{
		lastUpdate: time.Now(),
		rate:       r,
		burst:      b,
		tokens:     float64(b),
	}
}

func (l *limiter) allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()

	elapsed := now.Sub(l.lastUpdate).Seconds()
	l.tokens += elapsed * l.rate

	if l.tokens > float64(l.burst) {
		l.tokens = float64(l.burst)
	}

	l.lastUpdate = now

	if l.tokens >= 1 {
		l.tokens--
		return true
	}

	return false
}

func NewRateLimiter(r float64, b int) *VodkaRateLimiter {
	vrl := &VodkaRateLimiter{
		visitors: make(map[string]*visitor),
		rate:     r,
		burst:    b,
		done: make(chan struct{}),
	}

	go vrl.cleanup()
	return vrl
}

func (vrl *VodkaRateLimiter) cleanup() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	
	for {
		select{
		case <- ticker.C:
			
			vrl.mu.Lock()
			for ip, v := range vrl.visitors {
				if time.Since(v.lastSeen) > 5*time.Minute {
					delete(vrl.visitors, ip)
				}
			}
			vrl.mu.Unlock()

		case <- vrl.done:
			return
		}
	}
}

func (vrl *VodkaRateLimiter ) Stop(){
	vrl.once.Do(func() {
		close(vrl.done)
	})
}

func (vrl *VodkaRateLimiter) getVisitor(ip string) *limiter {
	vrl.mu.Lock()
	defer vrl.mu.Unlock()

	v, exists := vrl.visitors[ip]
	if !exists {
		l := newLimiter(vrl.rate, vrl.burst)
		vrl.visitors[ip] = &visitor{l, time.Now()}
		return l
	}

	v.lastSeen = time.Now()
	return v.l
}

func RateLimiter(vrl *VodkaRateLimiter) vodka.HandlerFunc {
	return func(c *vodka.Context) {
		ip := c.ClientIP()

		limiter := vrl.getVisitor(ip)

		if !limiter.allow() {
			c.Error(429, errors.New("rate limit exceeded"))
			return
		}
		
		c.Next()
	}
}
