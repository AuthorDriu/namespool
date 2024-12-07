package model

import (
	"fmt"

	dbTypes "github.com/AuthorDriu/namespool/internal/repository"
	db "github.com/AuthorDriu/namespool/internal/repository/sqlite"
)

func NewUser(nickname string, password []byte) (*dbTypes.User, error) {
	const op = "model.NewUser()"

	id, err := db.InsertUser(nickname, password)
	if err != nil {
		return nil, fmt.Errorf("%q: %v", op, err)
	}

	return &dbTypes.User{
		Id:       id,
		Nickname: nickname,
		Password: password,
	}, nil
}

func GetUser(nickname string) (*dbTypes.User, error) {
	const op = "model.GetUser()"

	user, err := db.SelectUser(nickname)
	if err != nil {
		return user, fmt.Errorf("%q: %v", op, err)
	}

	return user, nil
}
