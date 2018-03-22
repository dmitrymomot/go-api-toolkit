package auth

import "errors"

var (
	// ErrorUserNotFound is an error helper
	ErrorUserNotFound = errors.New("User not found")
	// ErrorInvalidToken is an error helper
	ErrorInvalidToken = errors.New("Invalid token")
	// ErrorEmptyHeader is an error helper
	ErrorEmptyHeader = errors.New("Auth header is empty")
	// ErrorInvalidHeader is an error helper
	ErrorInvalidHeader = errors.New("Invalid auth header")
	// ErrorCouldNotHandleToken is an error helper
	ErrorCouldNotHandleToken = errors.New("Could not handle this token")
	// ErrorTokenExpired is an error helper
	ErrorTokenExpired = errors.New("Token was expired")
	// ErrorTokenExpiredOrNotActive is an error helper
	ErrorTokenExpiredOrNotActive = errors.New("Token is either expired or not active yet")
)
