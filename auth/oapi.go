package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
)

// InjectUserInfoInCtx injects user info in context only if JWT presents in header.
func InjectUserInfoInCtx(jwtSigningKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" || auth == "Bearer" {
			return
		}
		token, claims, err := VerifyToken(c.Request, jwtSigningKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"data":    nil,
				"status":  http.StatusUnauthorized,
				"error":   true,
				"message": err.Error(),
			})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"data":    nil,
				"status":  http.StatusUnauthorized,
				"error":   true,
				"message": "unauthorized",
			})
			return
		}

		c.Set(authKey, claims)
		ctx := WithCtxToken(c.Request.Context(), token.Raw)
		c.Request = c.Request.WithContext(ctx)
	}
}

// OAPIMiddleware adds auth middleware to secure endpoints.
func OAPIMiddleware(jwtSigningKey string, spec *openapi3.T) gin.HandlerFunc {
	return middleware.OapiRequestValidatorWithOptions(spec,
		&middleware.Options{
			ErrorHandler: func(c *gin.Context, message string, statusCode int) {
				if strings.HasPrefix(message, "error in openapi3filter.SecurityRequirementsError: security requirements failed") {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"data":    nil,
						"status":  http.StatusUnauthorized,
						"error":   true,
						"message": "unauthorized",
					})
					return
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"data":    nil,
					"status":  http.StatusBadRequest,
					"error":   true,
					"message": message,
				})
			},
			Options: openapi3filter.Options{
				MultiError: false,
				AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
					return validateAuth(input, []byte(jwtSigningKey))
				},
			},
		})
}

func validateAuth(input *openapi3filter.AuthenticationInput, jwtSigningKey []byte) error {
	token, _, err := VerifyToken(input.RequestValidationInput.Request, jwtSigningKey)
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("unauthorized")
	}

	return nil
}
