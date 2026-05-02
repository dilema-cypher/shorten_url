package middleware

import (
	"net/http"
	"strings"

	"github.com/dilema-cypher/shorten_url/internal/gateways"
	"github.com/dilema-cypher/shorten_url/internal/utils"
)

var authClient *gateways.AuthMeClient

func InitAuthMiddleware(authURL string) {
	authClient = gateways.NewAuthMeClient(authURL)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")

		if strings.TrimSpace(bearerToken) == "" {
			w.WriteHeader(http.StatusUnauthorized)
			utils.SetErrorOnResponse(w, "API Token not provided")
			return
		}

		if !strings.HasPrefix(bearerToken, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			utils.SetErrorOnResponse(w, "Invalid API Token format")
			return
		}

		token := strings.Split(bearerToken, " ")[1]

		if strings.TrimSpace(token) == "" {
			w.WriteHeader(http.StatusUnauthorized)
			utils.SetErrorOnResponse(w, "Invalid API Token")
			return
		}

		if err := authClient.DoRequest(token); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			utils.SetErrorOnResponse(w, "Invalid API Token")
			return
		}

		next.ServeHTTP(w, r)
	})
}
