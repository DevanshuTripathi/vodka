package vodka

import "net/http"

func AllowCORS(origins []string) HandlerFunc {
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
			// Set the necessary headers for POST requests
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		}

		// The Critical Preflight Check
		if c.Request.Method == "OPTIONS" {
			if allow {
				c.Writer.WriteHeader(http.StatusNoContent) // Send 204 Success
			} else {
				c.Writer.WriteHeader(http.StatusForbidden) // Send 403 Forbidden
			}

			// Stop the request from going to the router
			c.Abort()
			return
		}

		c.Next()
	}
}
