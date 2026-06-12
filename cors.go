package vodka

import (
	"net/http"
	"strconv"
)

// CORSConfig holds configuration options for the CORS middleware
type CORSConfig struct {
	Origins []string
	MaxAge  int
}

func AllowCORS(origins []string, config ...CORSConfig) HandlerFunc {
	maxAge := 0
	if len(config) > 0 && config[0].MaxAge > 0 {
		maxAge = config[0].MaxAge
	}

	return func(c *Context) {
		origin := c.Request.Header.Get("Origin")
		allow := false

		// Check if the origin is in your allowed list
		for _, o := range origins {
			if o == "*" || o == origin {
				allow = true
				break
			}
		}

		if allow {
			// Set the necessary headers
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Methods",
				"GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers",
				"Content-Type, Authorization")
		}

		// The Critical Preflight Check
		if c.Request.Method == "OPTIONS" {
			if allow {
				// MaxAge only makes sense on preflight, set it here
				if maxAge > 0 {
					c.Writer.Header().Set("Access-Control-Max-Age",
						strconv.Itoa(maxAge))
				}
				c.Writer.WriteHeader(http.StatusNoContent) // 204
			} else {
				c.Writer.WriteHeader(http.StatusForbidden) // 403
			}
			c.Abort()
			return
		}

		c.Next()
	}
}
