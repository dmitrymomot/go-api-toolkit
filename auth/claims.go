package auth

import (
	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

type claims struct {
	jwt.StandardClaims
	TokenID    string                 `json:"id,omitempty"`
	ID         string                 `json:"user_id,omitempty"`
	Role       string                 `json:"user_role,omitempty"`
	Name       string                 `json:"user_name,omitempty"`
	Disabled   bool                   `json:"is_disabled,omitempty"`
	CustomData map[string]interface{} `json:"custom,omitempty"`
}

// IssuedTime returns token issued time
func (c *claims) IssuedTime() int64 {
	return c.IssuedAt
}

// ExpTime returns token issued time
func (c *claims) ExpTime() int64 {
	return c.ExpiresAt
}

// UserID returns user's id
func (c *claims) UserID() string {
	return c.ID
}

// UserRole returns user's role
func (c *claims) UserRole() string {
	return c.Role
}

// UserName returns user's role
func (c *claims) UserName() string {
	return c.Name
}

// IsDisabled returns whether user is disabled or not
func (c *claims) IsDisabled() bool {
	return c.Disabled
}

// Payload returns all token payload data
func (c *claims) Payload() map[string]interface{} {
	return c.CustomData
}
