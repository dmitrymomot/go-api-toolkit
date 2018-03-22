package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Default constants
const (
	ContentTypeHeader = "Content-Type"
	AcceptHeader      = "Accept"
	EmptyAcceptHeader = "*/*"
	GetRequestMethod  = "get"
)

// GinContentTypeMiddleware middleware for using with gin-gonic framework
func GinContentTypeMiddleware(c *gin.Context) {
	if !checkAcceptHeader(c.Request, config.ContentType()) {
		Err(c, http.StatusNotAcceptable, http.StatusText(http.StatusNotAcceptable), nil)
		return
	}
	if !checkContentTypeHeader(c.Request, config.ContentType()) {
		Err(c, http.StatusUnsupportedMediaType, http.StatusText(http.StatusUnsupportedMediaType), nil)
		return
	}
	c.Next()
}

func checkAcceptHeader(r *http.Request, need string) bool {
	h := strings.ToLower(r.Header.Get(AcceptHeader))
	return h != EmptyAcceptHeader && h == need
}

func checkContentTypeHeader(r *http.Request, need string) bool {
	if isGetMethod(r) {
		return true
	}
	h := strings.ToLower(r.Header.Get(ContentTypeHeader))
	return h == need
}

func isGetMethod(r *http.Request) bool {
	return strings.ToLower(r.Method) == GetRequestMethod
}
