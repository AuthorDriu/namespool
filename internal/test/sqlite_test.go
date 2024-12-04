package sqlite_test

import (
	"testing"

	repo "github.com/AuthorDriu/namespool/internal/repository/sqlite"
)

func TestPrepareDatabase(t *testing.T) {
	if err := repo.Prepare("store/test_database.db"); err != nil {
		t.Error(err)
	}
}
