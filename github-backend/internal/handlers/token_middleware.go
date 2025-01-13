package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

func TokenFromHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "Missing Authorization header"})
			return
		}

		// Validate and extract the token (Bearer <token>)
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "Invalid Authorization header format"})
			return
		}

		token := parts[1]

		// Pass token to the handler as a request header or context value
		ctx := context.WithValue(r.Context(), "accessToken", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
