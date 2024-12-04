package repository

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type database struct {
	conn *sql.DB
	mux  sync.Mutex
}

var db database

const schemaSQL = `
CREATE TABLE IF NOT EXISTS users (
	id        INT PRIMARY KEY,
	nickname  VARCHAR(50) UNIQUE NOT NULL,
	password  BLOB NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS user_nickname ON users(nickname); 

CREATE TABLE IF NOT EXISTS ideas (
	id INT PRIMARY KEY,
	title VARCHAR(250) NOT NULL,
	description TEXT,
	access INT NOT NULL,
	owner VARCHAR(50) NOT NULL,

	FOREIGN KEY (owner) REFERENCES users(nickname) 
);

CREATE INDEX IF NOT EXISTS idea_owner ON ideas(owner);
`

func Prepare(path string) error {
	const op string = "repositiry.sqlite.Prepare()"

	conn, err := sql.Open("sqlite3", path)
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
