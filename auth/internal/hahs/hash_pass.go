package hahs

import "golang.org/x/crypto/bcrypt"

func HashPass(pass string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(h), nil
}
