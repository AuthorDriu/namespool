package secure

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func getSecretKey() (string, error) {
	secretKey := os.Getenv("SECRETKEY")
	if secretKey == "" {
		return "", errors.New("SECRETKEY required in environ variables")
	}
	return secretKey, nil
}

func GenerateJWT(nickname string) (string, error) {
	const op = "secure.GenerateJWT()"

	payload := jwt.MapClaims{
		"nickname": nickname,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	secretKey, err := getSecretKey()
	if err != nil {
		return "", fmt.Errorf("%q: %v", op, err)
	}

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return tokenString, fmt.Errorf("%q: %v", op, err)
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (string, error) {
	const op = "secure.ParseJWT()"

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		secretKey, err := getSecretKey()
		if err != nil {
			return "", err
		}
		return secretKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("%q: %v", op, err)
	}

	if token == nil {
		return "", fmt.Errorf("%q: %v", op, errors.New("invalid token"))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("%q: %v", op, errors.New("cannot read MapClaims"))
	}

	exp, ok := claims["exp"].(int64)
	if !ok {
		return "", fmt.Errorf("%q: %v", op, errors.New("cannot read exp"))
	}

	if exp < time.Now().Unix() {
		return "", fmt.Errorf("%q: %v", op, errors.New("token expired"))
	}

	nickname, ok := claims["nickname"].(string)
	if !ok {
		return "", fmt.Errorf("%q: %v", op, errors.New("cannot read nickname"))
	}

	return nickname, nil
}
