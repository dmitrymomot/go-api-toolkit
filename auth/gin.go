// Package auth helpers for working with gin-gonic framework
package auth

import "github.com/gin-gonic/gin"

// UserID returns user id from gin context
func UserID(c *gin.Context) string {
	return contextString(c, JWTUserID)
}

// UserName returns user name from gin context
func UserName(c *gin.Context) string {
	return contextString(c, JWTUserName)
}

// UserNameDefault returns user name from gin context
func UserNameDefault(c *gin.Context, defVal string) string {
	return contextStringDefault(c, JWTUserName, defVal)
}

// UserRole returns user role from gin context
func UserRole(c *gin.Context) string {
	return contextString(c, JWTUserRole)
}

// UserRoleDefault returns user role from gin context
func UserRoleDefault(c *gin.Context, defVal string) string {
	return contextStringDefault(c, JWTUserRole, defVal)
}

func contextStringDefault(c *gin.Context, key, def string) string {
	result := contextString(c, key)
	if result == "" {
		return def
	}
	return result
}

func contextString(c *gin.Context, key string) string {
	val, ok := c.Get(key)
	if !ok || val == nil {
		return ""
	}
	return val.(string)
}
