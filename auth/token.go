package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

// Claims is part of jwt token.
type Claims struct {
	// User unique id
	UserID string `json:"userID"`

	UserName string `json:"userName"`
	Email    string `json:"email"`

	IsVerified bool   `json:"isVerified"`
	Role       string `json:"role"`

	jwt.StandardClaims
}

// GenerateToken for the given Claims.
func (claims *Claims) GenerateToken(jwtSigningKey []byte) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(jwtSigningKey)
}

// VerifyToken from Authorization header using signing key.
func VerifyToken(r *http.Request, jwtSigningKey []byte) (*jwt.Token, *Claims, error) {
	tokenString := r.Header.Get("Authorization")
	splitted := strings.Split(tokenString, " ")
	if len(splitted) != 2 {
		return nil, nil, errors.New("unexpected bearer token")
	}
	return verifyToken(splitted[1], jwtSigningKey)
}

func verifyToken(tokenString string, jwtSigningKey []byte) (*jwt.Token, *Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSigningKey, nil
	})
	return token, &claims, err
}
