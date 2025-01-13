package internalerrors

import "errors"

var ErrInternal error = errors.New("internal server error")
var ErrNotFound error = errors.New("not found")
var ErrUnauthorized error = errors.New("unauthorized")
var ErrCodeExpired error = errors.New("code expired")

func ProcessInternalError(err error) error {
	if errors.Is(err, ErrNotFound) {
		return ErrNotFound
	}
	return ErrInternal
}
