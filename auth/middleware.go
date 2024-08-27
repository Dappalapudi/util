package auth

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/Dappalapudi/util/env"
)

const authKey = "Auth"

// AnonymousRoutes allow bypassing of JWT verification for selected routes.
// Example: allow POST /auth/v1/login so user could login, because user initially has no JWT token.
type AnonymousRoutes struct {
	routes []*route
}

type route struct {
	method string
	path   string
}

// Add anonymous routes.
func (rt *AnonymousRoutes) Add(method, path string) {
	rt.routes = append(rt.routes, &route{method: method, path: path})
}

// ValidateToken checks jwt Authorization token. Use jwtSigningKey from secrets package.
func ValidateToken(jwtSigningKey []byte, anonRoutes *AnonymousRoutes) gin.HandlerFunc {
	return func(c *gin.Context) {

		if env.IsTest() {
			// no JWT validation in unit tests
			c.Next()
			return
		}

		if anonRoutes != nil {
			for _, route := range anonRoutes.routes {
				// Anonymous routes can match with wildcards.
				// "/home/catch/*", "/home/catch/foo"
				ok, err := filepath.Match(route.path, c.Request.URL.Path)
				if route.method == c.Request.Method && ok && err == nil {
					c.Next()
					return
				}
			}
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
		// Raw token is saved inside Context. It is used later to call other services
		// via /clients/ packages.
		ctx := WithCtxToken(c.Request.Context(), token.Raw)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// ValidateTokenDualMode is used for specific handlers that
// require serving in both logged in and not logged in modes.
// checks jwt Authorization token. Use jwtSigningKey from secrets package.
func ValidateTokenDualMode(jwtSigningKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {

		if env.IsTest() {
			// no JWT validation in unit tests
			c.Next()
			return
		}

		token, claims, err := VerifyToken(c.Request, jwtSigningKey)
		if err != nil {
			c.Next()
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
		// Raw token is saved inside Context. It is used later to call other services
		// via /clients/ packages.
		ctx := WithCtxToken(c.Request.Context(), token.Raw)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// GetClaims stored by middleware.
func GetClaims(c *gin.Context) *Claims {
	clm, _ := c.Get(authKey)
	if clm != nil {
		clm, ok := clm.(*Claims)
		if !ok {
			// this is not *Claims
			return nil
		}
		return clm
	}
	return nil
}
