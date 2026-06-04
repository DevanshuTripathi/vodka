package mixers

import (
    "net/http"
    "strconv"
    "sync"
    "time"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// PrometheusConfig holds configuration for the Prometheus middleware
type PrometheusConfig struct {
    MetricsPath  string
    Namespace    string
    ExcludePaths []string
}

// DefaultPrometheusConfig is the default configuration
var DefaultPrometheusConfig = PrometheusConfig{
    MetricsPath:  "/metrics",
    Namespace:    "vodka",
    ExcludePaths: []string{},
}

var (
    once             sync.Once
    requestsTotal    *prometheus.CounterVec
    requestDuration  *prometheus.HistogramVec
    activeRequests   prometheus.Gauge
    responseSizeBytes *prometheus.HistogramVec
    errorsTotal      *prometheus.CounterVec
)

func initMetrics(namespace string) {
    once.Do(func() {
        requestsTotal = prometheus.NewCounterVec(
            prometheus.CounterOpts{
                Namespace: namespace,
                Name:      "requests_total",
                Help:      "Total number of HTTP requests",
            },
            []string{"method", "path", "status"},
        )

        requestDuration = prometheus.NewHistogramVec(
            prometheus.HistogramOpts{
                Namespace: namespace,
                Name:      "request_duration_seconds",
                Help:      "HTTP request duration in seconds",
                Buckets:   prometheus.DefBuckets,
            },
            []string{"method", "path"},
        )

        activeRequests = prometheus.NewGauge(
            prometheus.GaugeOpts{
                Namespace: namespace,
                Name:      "active_requests",
                Help:      "Number of active HTTP requests",
            },
        )

        responseSizeBytes = prometheus.NewHistogramVec(
            prometheus.HistogramOpts{
                Namespace: namespace,
                Name:      "response_size_bytes",
                Help:      "HTTP response size in bytes",
                Buckets:   []float64{100, 1000, 10000, 100000},
            },
            []string{"method", "path"},
        )

        errorsTotal = prometheus.NewCounterVec(
            prometheus.CounterOpts{
                Namespace: namespace,
                Name:      "errors_total",
                Help:      "Total number of HTTP errors",
            },
            []string{"method", "path"},
        )

        prometheus.MustRegister(
            requestsTotal,
            requestDuration,
            activeRequests,
            responseSizeBytes,
            errorsTotal,
        )
    })
}

// Prometheus returns a middleware that collects Prometheus metrics
func Prometheus(config ...PrometheusConfig) func(http.Handler) http.Handler {
    cfg := DefaultPrometheusConfig
    if len(config) > 0 {
        cfg = config[0]
        if cfg.MetricsPath == "" {
            cfg.MetricsPath = DefaultPrometheusConfig.MetricsPath
        }
        if cfg.Namespace == "" {
            cfg.Namespace = DefaultPrometheusConfig.Namespace
        }
    }

    initMetrics(cfg.Namespace)

    excludeMap := make(map[string]bool)
    for _, p := range cfg.ExcludePaths {
        excludeMap[p] = true
    }

    return func(next http.Handler) http.Handler {
        // Expose /metrics endpoint
        http.Handle(cfg.MetricsPath, promhttp.Handler())

        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Skip excluded paths
            if excludeMap[r.URL.Path] {
                next.ServeHTTP(w, r)
                return
            }

            start := time.Now()
            activeRequests.Inc()

            // Wrap response writer to capture status code and size
            wrapped := &responseWriter{ResponseWriter: w, status: 200}
            next.ServeHTTP(wrapped, r)

            duration := time.Since(start).Seconds()
            status := strconv.Itoa(wrapped.status)
            path := r.URL.Path
            method := r.Method

            activeRequests.Dec()
            requestsTotal.WithLabelValues(method, path, status).Inc()
            requestDuration.WithLabelValues(method, path).Observe(duration)
            responseSizeBytes.WithLabelValues(method, path).Observe(
                float64(wrapped.size),
            )

            if wrapped.status >= 500 {
                errorsTotal.WithLabelValues(method, path).Inc()
            }
        })
    }
}

// responseWriter wraps http.ResponseWriter to capture status and size
type responseWriter struct {
    http.ResponseWriter
    status int
    size   int
}

func (rw *responseWriter) WriteHeader(status int) {
    rw.status = status
    rw.ResponseWriter.WriteHeader(status)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
    size, err := rw.ResponseWriter.Write(b)
    rw.size += size
    return size, err
}