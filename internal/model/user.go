package model

import "errors"

type User struct {
	Id       int
	Nickname string // unique nickname
	Password []byte // hashed password
}

var ErrNotUniqueUsername = errors.New("not unique user")

func NewUser(nickname string, password []byte) (*User, error) {

	newUser := &User{
		Nickname: nickname,
		Password: password,
	}

	return newUser, errors.New("not implemented")
}
