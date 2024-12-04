package model

import "errors"

type AccessModifier int

const (
	Public  AccessModifier = 1
	Private AccessModifier = 0
)

type Idea struct {
	Id          int
	Title       string
	Description string
	Access      AccessModifier
	Owner       string // users.nickname in db
}

func NewNameIdia(title string, description string, owner string) (*Idea, error) {
	return nil, errors.New("not implemented")
}
