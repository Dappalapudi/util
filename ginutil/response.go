package ginutil

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Dappalapudi/util/auth"
)

// JSON sends StatusOK, formatted message and data.
//
// Example 1: ginutil.JSON(c, &userInfo, "login success")
// Example 2: ginutil.JSON(c, &userDetails, "Email has been sent to: %s", email)
// Example 3: ginutil.JSON(c, &details, "Email sent to: %s, token will expire in %d minutes", email, expires)
func JSON(c *gin.Context, data interface{}, format string, v ...interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf(format, v...),
		"error":   false,
		"data":    data,
	})
}

// JSONError sends passed in http status code, formatted message and data.
//
// Example 1: ginutil.JSONError(c, http.StatusNotFound, &userInfo, "login failed, no such user")
func JSONError(c *gin.Context, status int, data interface{}, format string, v ...interface{}) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": fmt.Sprintf(format, v...),
		"error":   true,
		"data":    data,
	})
}

// UserIDFromJWT returns user id from JWT claims, if it can't find claims returns empty string.
// It expects that it's used under auth middleware.
func UserIDFromJWT(c *gin.Context) string {
	claims := auth.GetClaims(c)
	if claims == nil {
		return ""
	}
	return claims.UserID
}
