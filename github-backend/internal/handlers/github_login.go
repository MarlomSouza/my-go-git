package handlers

import (
	"net/http"
	"time"

	internalerrors "github.com/MarlomSouza/go-git/internal-errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang.org/x/oauth2"
)

// OAuthHandler handles GitHub OAuth operations
type OAuthHandler struct {
	OAuthConfig         *oauth2.Config
	FrontendRedirectURL string
}

func (h *OAuthHandler) RegisterRoutes(r chi.Router) {

	r.Get("/login/github", h.GitHubLogin)
	r.Get("/login/github/callback", h.GitHubCallback)
	r.Post("/logout", h.Logout)
}

// NewOAuthHandler creates a new instance of OAuthHandler
func NewOAuthHandler(config *oauth2.Config, frontendRedirectURL string) *OAuthHandler {
	return &OAuthHandler{OAuthConfig: config, FrontendRedirectURL: frontendRedirectURL}
}

// Handle GitHub Login
func (h *OAuthHandler) GitHubLogin(w http.ResponseWriter, r *http.Request) {
	url := h.OAuthConfig.AuthCodeURL("random_state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

// GitHub Callback Handler
func (h *OAuthHandler) GitHubCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		render.JSON(w, r, internalerrors.ErrInternal)
		return
	}

	// Exchange the code for an access token
	token, err := h.OAuthConfig.Exchange(r.Context(), code)
	if err != nil {
		render.JSON(w, r, internalerrors.ErrInternal)
		return
	}

	// Store the token in a secure HTTP-only cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token.AccessToken,
		Path:     "/",
		HttpOnly: false,
		Secure:   true,
		Expires:  token.Expiry,
	})

	http.Redirect(w, r, h.FrontendRedirectURL, http.StatusFound)

}

// Handle logout by clearing the cookie
func (h *OAuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0), // Expire immediately
		HttpOnly: false,
		Secure:   true,
	})

	render.JSON(w, r, "Logged out successfully")
}
