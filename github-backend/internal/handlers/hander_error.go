package handlers

import (
	"errors"
	"net/http"
	"time"

	internalerrors "github.com/MarlomSouza/go-git/internal-errors"
	"github.com/go-chi/render"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(endpointFunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, statusCode, err := endpointFunc(w, r)
		if err != nil {
			if errors.Is(err, internalerrors.ErrInternal) {
				render.Status(r, http.StatusInternalServerError)
			} else if errors.Is(err, internalerrors.ErrNotFound) {
				render.Status(r, http.StatusNotFound)
			} else if errors.Is(err, internalerrors.ErrUnauthorized) {

				http.SetCookie(w, &http.Cookie{
					Name:     "access_token",
					Value:    "",
					Path:     "/",
					Expires:  time.Unix(0, 0), // Expire immediately
					HttpOnly: false,
					Secure:   true,
				})

				render.Status(r, http.StatusUnauthorized)
			} else {
				render.Status(r, http.StatusBadRequest)
			}

			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, statusCode)

		if obj == nil && statusCode == http.StatusNoContent {
			render.JSON(w, r, nil)
		}

		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
