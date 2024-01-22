package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key") // Change this to a secure secret key

// Authenticate is a middleware to enforce JWT authentication.
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only enforce authentication for the /reports endpoint
		if strings.HasPrefix(r.URL.Path, "/reports") {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			claims := &jwt.StandardClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return secretKey, nil
			})

			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		// Continue to the next handler if authentication is successful
		next.ServeHTTP(w, r)
	})
}
