package mixers

import (
	"runtime"
	"testing"
	"time"
)

func TestLimiterAllow(t *testing.T) {
	l := newLimiter(1, 2)

	if !l.allow() {
		t.Fatal("expected first request to pass")
	}

	if !l.allow() {
		t.Fatal("expected second request to pass")
	}

	if l.allow() {
		t.Fatal("expected third request to fail")
	}
}

func TestLimiterRefill(t *testing.T) {
	l := newLimiter(1, 1)

	if !l.allow() {
		t.Fatal("expected first request to pass")
	}

	if l.allow() {
		t.Fatal("expected second request to fail")
	}

	time.Sleep(1100 * time.Millisecond)

	if !l.allow() {
		t.Fatal("expected token refill")
	}
}

func TestRateLimiterStop(t *testing.T) {
	before := runtime.NumGoroutine()

	vrl := NewRateLimiter(1, 5)

	time.Sleep(100 * time.Millisecond)

	vrl.Stop()

	time.Sleep(200 * time.Millisecond)

	after := runtime.NumGoroutine()

	if after > before+2 {
		t.Fatalf(
			"possible goroutine leak detected: before=%d after=%d",
			before,
			after,
		)
	}
}
