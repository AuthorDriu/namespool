package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/AuthorDriu/namespool/pkg/path"

	_ "github.com/mattn/go-sqlite3"
)

type database struct {
	conn *sql.DB
	mux  sync.Mutex
}

var db database

const schemaSQL = `
CREATE TABLE IF NOT EXISTS users (
	id        INTEGER PRIMARY KEY AUTOINCREMENT,
	nickname  VARCHAR(50) UNIQUE NOT NULL,
	password  BLOB NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS user_nickname ON users(nickname); 

CREATE TABLE IF NOT EXISTS ideas (
	id           INTEGER PRIMARY KEY AUTOINCREMENT,
	title        VARCHAR(250) NOT NULL,
	description  TEXT,
	access       INTEGER NOT NULL,
	owner        VARCHAR(50) NOT NULL,

	FOREIGN KEY (owner) REFERENCES users(nickname) 
);

CREATE INDEX IF NOT EXISTS idea_owner ON ideas(owner);
CREATE UNIQUE INDEX IF NOT EXISTS users_idea ON ideas(title, owner);`

func Prepare(pathToDatabaseFile string) error {
	const op = "repositiry.sqlite.Prepare()"
	pathToDatabaseFile = path.FromRoot(pathToDatabaseFile)

	if _, err := os.Stat(pathToDatabaseFile); errors.Is(err, os.ErrNotExist) {
		if file, err := os.Create(pathToDatabaseFile); err != nil {
			return fmt.Errorf("%q: %v", op, err)
		} else {
			file.Close()
		}
	}

	conn, err := sql.Open("sqlite3", pathToDatabaseFile)
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	_, err = conn.Exec(schemaSQL)
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	db = database{conn: conn}
	return nil
}

func Close() error {
	const op = "repository.sqlite.Close()"
	if err := db.conn.Close(); err != nil {
		return fmt.Errorf("%q: %v", op, err)
	}
	return nil
}
