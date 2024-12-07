package model

import (
	"fmt"

	dbTypes "github.com/AuthorDriu/namespool/internal/repository"
	db "github.com/AuthorDriu/namespool/internal/repository/sqlite"
)

func NewIdia(title string, description string, access dbTypes.AccessModifier, owner string) (*dbTypes.Idea, error) {
	const op = "model.NewIdea()"

	id, err := db.InsertIdea(title, description, access, owner)
	if err != nil {
		return nil, fmt.Errorf("%q: %v", op, err)
	}

	return &dbTypes.Idea{
		Id:          id,
		Title:       title,
		Description: description,
		Access:      access,
		Owner:       owner,
	}, nil
}

func DeleteIdea(owner string, title string) error {
	const op = "model.DeleteIdea()"

	err := db.DeleteIdea(owner, title)
	if err != nil {
		return fmt.Errorf("%q: %v", op, err)
	}

	return nil
}

func GetIdea(owner string, title string) (*dbTypes.Idea, error) {
	const op = "model.GetIdea()"

	idea, err := db.SelectIdea(owner, title)
	if err != nil {
		return idea, fmt.Errorf("%q: %v", op, err)
	}

	return idea, nil
}

func GetIdeasByUser(owner string) ([]*dbTypes.Idea, error) {
	const op = "model.GetIdeasByUser()"

	ideas, err := db.SelectIdeasByUser(owner)
	if err != nil {
		return ideas, fmt.Errorf("%q: %v", op, err)
	}

	return ideas, nil
}

func GetPublicIdeasByUser(owner string) ([]*dbTypes.Idea, error) {
	const op = "model.GetIdeasByUser()"

	ideas, err := db.SelectPublicIdeasByUser(owner)
	if err != nil {
		return ideas, fmt.Errorf("%q: %v", op, err)
	}

	return ideas, nil
}
