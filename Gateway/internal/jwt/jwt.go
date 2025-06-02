package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/bobr-lord-messenger/gateway/internal/jwtutil"
	"strings"
)

func ParseJWT(tokenStr string) (string, error) {
	parts := strings.SplitN(tokenStr, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid token")
	}
	tokenStr = parts[1]

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return jwtutil.GetPublicKey()
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}

	userID, ok := claims["user_id"]
	if !ok {
		return "", fmt.Errorf("user_id not found in token")
	}

	return userID.(string), nil
}
