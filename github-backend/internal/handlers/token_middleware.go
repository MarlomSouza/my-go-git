package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
)

func TokenFromCookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from the "access_token" cookie
		cookie, err := r.Cookie("access_token")
		if err != nil || cookie == nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "Missing cookie"})
			return
		}

		token := cookie.Value

		// Pass token to the handler as a request context value
		ctx := context.WithValue(r.Context(), "accessToken", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
