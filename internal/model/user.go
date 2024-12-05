package model

import (
	"errors"
	"fmt"

	"github.com/AuthorDriu/namespool/internal/repository"
	"github.com/AuthorDriu/namespool/internal/repository/sqlite"
)

var ErrNotUniqueUsername = errors.New("not unique user")

func NewUser(nickname string, password []byte) (*repository.User, error) {
	const op = "model.NewUser()"

	id, err := sqlite.InsertUser(nickname, password)
	if err != nil {
		return nil, fmt.Errorf("%q: %v", op, err)
	}

	return &repository.User{
		Id:       id,
		Nickname: nickname,
		Password: password,
	}, nil
}
