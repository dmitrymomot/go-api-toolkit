package auth

import "time"

const (
	// authHeader is the default name of header key with token
	authHeader = "Authorization"
	// authTokenType is the default JWT type
	authTokenType = "Bearer"
	// tokenMaxAge is default token max age in seconds
	tokenMaxAge = 3600
)

// Configer config interface
type Configer interface {
	TokenSigningKey() []byte
	AuthHeaderKey() string
	AuthTokenType() string
	TokenLifetime() time.Duration
	TokenRefreshTime() time.Duration
}

// Config auth package configuration
type Config struct {
	SigningKey      string
	TokenType       string
	AuthHeader      string
	TokenMaxAge     int64
	TokenMaxRefresh int64
}

// TokenSigningKey returns signing key
func (c *Config) TokenSigningKey() []byte {
	return []byte(c.SigningKey)
}

// AuthHeaderKey returns the auth header
func (c *Config) AuthHeaderKey() string {
	if c.AuthHeader == "" {
		c.AuthHeader = authHeader
	}
	return c.AuthHeader
}

// AuthTokenType returns the auth header
func (c *Config) AuthTokenType() string {
	if c.TokenType == "" {
		c.TokenType = authTokenType
	}
	return c.TokenType
}

// TokenLifetime returns token lifetime
func (c *Config) TokenLifetime() time.Duration {
	if c.TokenMaxAge <= 0 {
		c.TokenMaxAge = tokenMaxAge
	}
	return time.Duration(c.TokenMaxAge) * time.Second
}

// TokenRefreshTime returns token lifetime
func (c *Config) TokenRefreshTime() time.Duration {
	return time.Duration(c.TokenMaxRefresh) * time.Second
}
