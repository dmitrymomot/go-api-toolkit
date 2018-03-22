package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

// JWT struct
type jsonWebToken struct {
	config Configer
}

// GenToken generates json web token
func (j *jsonWebToken) GenToken(user Userer) (Tokener, error) {
	cl := &claims{}
	cl.ExpiresAt = time.Now().Add(j.config.TokenLifetime()).Unix()
	cl.IssuedAt = user.ResetTokenTime()
	cl.ID = user.PrimaryKey()
	cl.Role = user.BaseRole()
	cl.Disabled = user.IsDisabled()
	cl.CustomData = user.Payload()
	cl.TokenID = uuid.Must(uuid.NewV1()).String()

	t := jwt.New(jwt.SigningMethodHS256)
	t.Claims = cl

	ss, err := t.SignedString(j.config.TokenSigningKey())
	if err != nil {
		return nil, err
	}
	token := &token{
		Expire: time.Unix(cl.ExpTime(), 0),
		Token:  ss,
		Type:   j.config.AuthTokenType(),
	}
	return token, nil
}

// CheckToken checks whether a token is valid
func (j *jsonWebToken) ParseFromHeader(headers Headerer) (Claimer, error) {
	token, err := j.parseTokenFromHeader(headers)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrorInvalidToken
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, ErrorTokenExpiredOrNotActive
			} else {
				return nil, ErrorCouldNotHandleToken
			}
		}
		return nil, ErrorInvalidToken
	}
	claims, err := j.parseClaimsFromTokenString(token)
	if err != nil {
		return nil, ErrorInvalidToken
	}
	return claims, nil
}

// ParseTokenFromHeader returns token string from request headers
func (j *jsonWebToken) parseTokenFromHeader(headers Headerer) (string, error) {
	authHeader := headers.Get(j.config.AuthHeaderKey())
	if authHeader == "" {
		return "", ErrorEmptyHeader
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == j.config.AuthTokenType()) {
		return "", ErrorInvalidHeader
	}
	return parts[1], nil
}

// ParseClaimsFromTokenString helper
func (j *jsonWebToken) parseClaimsFromTokenString(tokenString string) (Claimer, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.config.TokenSigningKey(), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*claims)
	if !ok || !token.Valid {
		return nil, ErrorCouldNotHandleToken
	}
	return claims, nil
}

// HTTPMiddleware of JWT
func (j *jsonWebToken) HTTPMiddleware(request *http.Request, fetchUser FetchUserByIDFunc) (Claimer, error) {
	claims, err := j.ParseFromHeader(request.Header)
	if err != nil {
		return nil, err
	}
	user, err := fetchUser(claims.UserID())
	if err != nil {
		return nil, ErrorUserNotFound
	}
	if user.ResetTokenTime() > claims.IssuedTime() {
		return nil, ErrorTokenExpired
	}
	return claims, nil
}

// Default gin context keys
const (
	JWTUserID     = "JWT_USER_ID"
	JWTUserRole   = "JWT_USER_ROLE"
	JWTUserName   = "JWT_USER_NAME"
	JWTCustomData = "JWT_CUSTOM_DATA"
)

// GinMiddleware middleware for using with gin-gonic framework
func (j *jsonWebToken) GinMiddleware(fetchUser FetchUserByIDFunc, errHandler func(c *gin.Context, err error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := j.HTTPMiddleware(c.Request, fetchUser)
		if err != nil {
			errHandler(c, err)
			return
		}
		c.Set(JWTUserID, claims.UserID())
		c.Set(JWTUserRole, claims.UserRole())
		c.Set(JWTUserName, claims.UserName())
		c.Set(JWTCustomData, claims.Payload())
		c.Next()
		return
	}
}
