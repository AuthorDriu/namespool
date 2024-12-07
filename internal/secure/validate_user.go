package secure

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EnctyptPassord(password string) ([]byte, error) {
	if len(password) > 72 {
		return nil, fmt.Errorf("too long password")
	}

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ValidatePassword(password string, passwordEnctypted []byte) bool {
	return bcrypt.CompareHashAndPassword(passwordEnctypted, []byte(password)) != nil
}
