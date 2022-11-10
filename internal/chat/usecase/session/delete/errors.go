package delete

import "errors"

var (
	ErrSessionNotFound = errors.New("session is not found")
	ErrSessionDeleted  = errors.New("session is deleted")
)
