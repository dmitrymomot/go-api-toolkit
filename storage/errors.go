package storage

import "errors"

// Default errors
var (
	ErrorConnection     = errors.New("Could not connect to storage")
	ErrorFileSaving     = errors.New("Could not save file into storage")
	ErrorFileDeleting   = errors.New("Could not delete file from storage")
	ErrorFileFetching   = errors.New("Could not get file from storage")
	ErrorURLFetching    = errors.New("Could not get file path")
	ErrorCouldNotUpload = errors.New("Could not upload file to storage")
)
