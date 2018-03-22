package api

import "strings"

// DefaultContentType constant
const DefaultContentType = "application/json"

// Configer config interface
type Configer interface {
	ContentType() string
}

// Config of the package
type Config struct {
	AllowedContentType string
}

// ContentType returns allowed to use content type
func (c *Config) ContentType() string {
	if c.AllowedContentType == "" {
		c.AllowedContentType = DefaultContentType
	}
	return strings.ToLower(c.AllowedContentType)
}
