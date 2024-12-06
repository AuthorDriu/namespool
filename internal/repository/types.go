package repository

type User struct {
	Id       int
	Nickname string // unique nickname
	Password []byte // hashed password
}

type AccessModifier int

const (
	IdeaPublic  AccessModifier = 1
	IdeaPrivate AccessModifier = 0
)

type Idea struct {
	Id          int
	Title       string
	Description string
	Access      AccessModifier
	Owner       string // users.nickname in db
}
