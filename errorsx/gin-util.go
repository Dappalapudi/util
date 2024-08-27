package errorsx

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrBadRequest   BadRequestError
	ErrNotFound     NotFoundError
	ErrUnauthorized UnauthorizedError
	ErrConflict     ConflictError
)

func HandleError(c *gin.Context, err error) {

	if errors.As(err, &ErrBadRequest) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if errors.As(err, &ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	if errors.As(err, &ErrUnauthorized) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}

	if errors.As(err, &ErrConflict) {
		c.JSON(http.StatusConflict, gin.H{
			"error":   true,
			"status":  http.StatusConflict,
			"message": err.Error(),
		})
		return
	}

	// Default to internal server error in all other cases.
	c.JSON(http.StatusInternalServerError, gin.H{
		"error":   true,
		"status":  http.StatusInternalServerError,
		"message": err.Error(),
	})

}
