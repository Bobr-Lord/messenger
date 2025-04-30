package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/bobr-lord-messenger/auth/internal/jwtutil"
	"time"
)

func CreateJWT(userID string) (string, error) {
	key, err := jwtutil.GetPrivateKey()
	if err != nil {
		return "", fmt.Errorf("Error creating JWT key: %v", err)
	}
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(key)
}
