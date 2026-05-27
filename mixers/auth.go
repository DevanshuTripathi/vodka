package mixers

import (
	"errors"
	"strings"

	"github.com/DevanshuTripathi/vodka"
)

// Token Validator type recieves token and returns the data linked to it and if the token is valid
type TokenValidator func(token string) (data any, isValid bool)

// Bearer Auth for getting the bearer token from the request header
// Then validate the token using the validator function
// Set value in the context for later use from a middleware
func BearerAuth(ctxKey string, validator TokenValidator) vodka.HandlerFunc {
	return func(c *vodka.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		// c.Error() routes through ErrorHandler and calls Abort internally
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.Error(401, errors.New("Unauthorized: Missing or malformed token"))
			return
		}

		providedToken := strings.TrimPrefix(authHeader, "Bearer ")

		data, isValid := validator(providedToken)

		if !isValid {
			c.Error(401, errors.New("Unauthorized: Invalid token"))
			return
		}

		c.Set(ctxKey, data)

		c.Next()
	}
}

