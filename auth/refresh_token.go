package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type RefreshClaims struct {
	jwt.StandardClaims
}

func NewRefreshClaims(userID string, ttl time.Duration, jwtIssuer string) RefreshClaims {
	return RefreshClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(ttl).Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
			Issuer:    jwtIssuer,
			Subject:   userID,
		},
	}
}

func (c RefreshClaims) GenerateToken(jwtSigningKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(jwtSigningKey)
}

func VerifyRefreshToken(tokenString string, jwtSigningKey []byte) (*jwt.Token, *RefreshClaims, error) {
	var claims RefreshClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSigningKey, nil
	})
	return token, &claims, err
}
