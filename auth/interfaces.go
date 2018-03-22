package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

// Tokener is a token interface
type Tokener interface {
	String() string
	MarshalJSON() ([]byte, error)
}

// Claimer claims struct interface
type Claimer interface {
	jwt.Claims
	IssuedTime() int64
	ExpTime() int64
	UserID() string
	UserRole() string
	UserName() string
	IsDisabled() bool
	Payload() map[string]interface{}
}

// Userer is a user model interface
type Userer interface {
	PrimaryKey() string
	BaseRole() string
	IsDisabled() bool
	ResetTokenTime() int64
	Payload() map[string]interface{}
	Name() string
}

// FetchUserByIDFunc function interface which fetches user by id
type FetchUserByIDFunc func(id string) (Userer, error)

// Headerer interface
type Headerer interface {
	Get(key string) string
}

// JWT json web token struct interface
type JWT interface {
	GenToken(user Userer) (Tokener, error)
	ParseFromHeader(headers Headerer) (Claimer, error)
	HTTPMiddleware(request *http.Request, fetchUser FetchUserByIDFunc) (Claimer, error)
	GinMiddleware(fetchUser FetchUserByIDFunc, errHandler func(c *gin.Context, err error)) gin.HandlerFunc
}
