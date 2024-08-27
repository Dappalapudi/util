package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRefreshClaimsToken(t *testing.T) {
	var signingKey = []byte("123")

	claims := NewRefreshClaims("userID", time.Hour, "test")

	tokenString, err := claims.GenerateToken(signingKey)
	require.NoError(t, err)

	_, c, err := VerifyRefreshToken(tokenString, signingKey)
	if err != nil {
		t.Errorf("verify token failed: %s", err)
		return
	}

	assert.Equal(t, claims, *c)
}
