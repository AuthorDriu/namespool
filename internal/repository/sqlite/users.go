package sqlite

import (
	"fmt"

	"github.com/AuthorDriu/namespool/internal/repository"

	_ "github.com/mattn/go-sqlite3"
)

func InsertUser(nickname string, password []byte) (int, error) {
	const op = "repositiry.sqlite.InsertUser()"

	stmt := `
	INSERT INTO users (nickname, password)
	VALUES (?, ?);`

	db.mux.Lock()
	result, err := db.conn.Exec(stmt, nickname, password)
	db.mux.Unlock()
	if err != nil {
		return -1, fmt.Errorf("%q: %v", op, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("%q: %v", op, err)
	}

	return int(id), nil
}

func SelectUser(nickname string) (*repository.User, error) {
	const op = "repositiry.sqlite.SelectUser()"

	query := `
	SELECT * FROM users
	WHERE nickname = ?;`

	db.mux.Lock()
	row := db.conn.QueryRow(query, nickname)
	db.mux.Unlock()

	var (
		id        int64
		_nickname string
		password  []byte
	)
	err := row.Scan(&id, &_nickname, &password)
	if err != nil {
		return nil, fmt.Errorf("%q: %v", op, err)
	}

	return &repository.User{
		Id:       int(id),
		Nickname: _nickname,
		Password: password,
	}, nil
}
