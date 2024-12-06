package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/AuthorDriu/namespool/internal/repository"
)

func InsertIdea(title, description string, access repository.AccessModifier, owner string) (int, error) {
	const op = "repository.sqlite.InsertIdea()"

	stmt := `
	INSERT INTO ideas (title, description, access, owner)
	VALUES (?, ?, ?, ?);`

	db.mux.Lock()
	result, err := db.conn.Exec(stmt, title, description, int(access), owner)
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

func DeleteIdea(owner string, title string) error {
	const op = "repository.sqlite.DeleteIdea()"

	stmt := `
	DELETE FROM ideas
	WHERE owner = ? AND title = ?;`

	db.mux.Lock()
	_, err := db.conn.Exec(stmt, owner, title)
	db.mux.Unlock()

	if err != nil {
		return fmt.Errorf("%q: %v", op, err)
	}

	return nil
}

func SelectIdea(owner string, title string) (*repository.Idea, error) {
	const op = "repository.sqlite.SelectIdea()"

	query := `
	SELECT * FROM ideas
	WHERE owner = ? AND title = ?;`

	db.mux.Lock()
	row := db.conn.QueryRow(query, owner, title)
	db.mux.Unlock()

	var (
		idea_id          int64
		idea_title       string
		idea_description string
		idea_access      int
		idea_owner       string
	)

	err := row.Scan(&idea_id, &idea_title, &idea_description, &idea_access, &idea_owner)

	if err != nil {
		return nil, fmt.Errorf("%q: %v", op, err)
	}

	return &repository.Idea{
		Id:          int(idea_id),
		Title:       idea_title,
		Description: idea_description,
		Access:      repository.AccessModifier(idea_access),
		Owner:       idea_owner,
	}, nil
}

func extractIdeas(rows **sql.Rows) ([]*repository.Idea, error) {
	ideas := make([]*repository.Idea, 0)
	for (*rows).Next() {
		var (
			idea_id          int64
			idea_title       string
			idea_description string
			idea_access      int64
			idea_owner       string
		)

		err := (*rows).Scan(&idea_id, &idea_title, &idea_description, &idea_access, &idea_owner)

		if err != nil {
			return ideas, err
		}

		idea := &repository.Idea{
			Id:          int(idea_id),
			Title:       idea_title,
			Description: idea_description,
			Access:      repository.AccessModifier(idea_access),
			Owner:       idea_owner,
		}

		ideas = append(ideas, idea)
	}
	return ideas, nil
}

func SelectIdeasByUser(owner string) ([]*repository.Idea, error) {
	const op = "repository.sqlite.SelectIdeasByUser()"

	query := `
	SELECT * FROM ideas
	WHERE owner = ?;`

	db.mux.Lock()
	rows, err := db.conn.Query(query, owner)
	db.mux.Unlock()

	if err != nil {
		return nil, fmt.Errorf("%q: %v", op, err)
	}
	defer rows.Close()

	ideas, err := extractIdeas(&rows)

	if err != nil {
		return ideas, fmt.Errorf("%q: %v", op, err)
	}

	return ideas, nil
}

func SelectPublicIdeasByUser(owner string) ([]*repository.Idea, error) {
	const op = "repository.sqlite.SelectIdeasByUser()"

	query := `
	SELECT * FROM ideas
	WHERE owner = ? AND access = 1;`

	db.mux.Lock()
	rows, err := db.conn.Query(query, owner)
	db.mux.Unlock()

	if err != nil {
		return nil, fmt.Errorf("%q: %v", op, err)
	}
	defer rows.Close()

	ideas, err := extractIdeas(&rows)

	if err != nil {
		return ideas, fmt.Errorf("%q: %v", op, err)
	}

	return ideas, nil
}
