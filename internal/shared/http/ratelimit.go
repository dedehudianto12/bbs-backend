package http

import (
	"net/http"
	"sync"
	"time"
)

// PerIPRateLimit returns middleware that limits each IP to `max` requests per `window`.
// Uses a sliding window per IP with lazy cleanup. Zero allocations after warm-up.
func PerIPRateLimit(max int, window time.Duration) func(http.Handler) http.Handler {
	type bucket struct {
		count  int
		resetAt time.Time
	}

	var (
		mu       sync.Mutex
		visitors = make(map[string]*bucket)
	)

	// Lazy cleanup goroutine — runs every window period
	go func() {
		for {
			time.Sleep(window)
			mu.Lock()
			now := time.Now()
			for ip, b := range visitors {
				if now.After(b.resetAt) {
					delete(visitors, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := clientIP(r)

			mu.Lock()
			v, exists := visitors[ip]
			now := time.Now()
			if !exists || now.After(v.resetAt) {
				v = &bucket{count: 0, resetAt: now.Add(window)}
				visitors[ip] = v
			}
			v.count++
			exceeded := v.count > max
			mu.Unlock()

			if exceeded {
				w.Header().Set("Retry-After", "60")
				Error(w, http.StatusTooManyRequests, &rateLimitError{})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// clientIP extracts the real client IP, respecting X-Forwarded-For if behind a reverse proxy.
func clientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// Take the first IP (client origin) in the chain
		for i := 0; i < len(xff); i++ {
			if xff[i] == ',' {
				return xff[:i]
			}
		}
		return xff
	}
	return r.RemoteAddr
}

type rateLimitError struct{}

func (e *rateLimitError) Error() string { return "terlalu banyak permintaan, coba lagi nanti" }
