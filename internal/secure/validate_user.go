package secure

import (
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

func EnctyptPassord(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ValidatePassword(password string, passwordEnctypted []byte) bool {
	encrypted, err := EnctyptPassord(password)
	if err != nil {
		return false
	}
	if !reflect.DeepEqual(encrypted, passwordEnctypted) {
		return false
	}
	return true
}
