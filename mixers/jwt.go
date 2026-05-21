package mixers

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// Validates a token on a secretKey and returns the claims
func JWTValidator(secretKey string) TokenValidator {
	return func(tokenString string) (any, bool) {
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			return nil, false
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			return claims, true
		}

		return nil, false
	}
}

// Creates a claims map for the key and payload
func GenerateJWT(secretKey string, payload map[string]any, expiresIn time.Duration) (string, error) {
	claims := jwt.MapClaims{}

	for key, value := range payload {
		claims[key] = value
	}

	claims["exp"] = time.Now().Add(expiresIn).Unix()
	claims["iat"] = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
