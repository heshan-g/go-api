package customMiddleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/heshan-g/go-api/config"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			msg := "No Authorization header provided"
			http.Error(w, msg, http.StatusUnauthorized)
			return
		}

		prefix, tokenString, found := strings.Cut(authHeader, " ")
		if !found || prefix != "Bearer" {
			msg := "Authorization header is in an invalid format"
			http.Error(w, msg, http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method (%v)", token.Header["alg"])
			}

			return config.LoadPublicKey()
		})
		if err != nil {
			msg := fmt.Sprintf("Bad token: %s", err.Error())
			http.Error(w, msg, http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			msg := "Unexpected error (error mapping token claims)"
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		// REMOVE BELOW!
		for key, value := range claims {
			fmt.Println(key,": ", value)
		}

		next.ServeHTTP(w, r)
	})
}
