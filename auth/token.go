package auth

import (
	"encoding/json"
	"time"
)

type token struct {
	Expire time.Time `json:"expire"`
	Token  string    `json:"token"`
	Type   string    `json:"type"`
}

// String returns token string
func (t *token) String() string {
	return t.Token
}

// MarshalJSON custom json marshaling
func (t *token) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}
