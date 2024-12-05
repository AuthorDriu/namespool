package sqlite_test

import (
	"os"
	"reflect"
	"testing"

	repo "github.com/AuthorDriu/namespool/internal/repository/sqlite"
	"github.com/AuthorDriu/namespool/pkg/path"
)

const databasePath string = "store/test_database.db"

func TestPrepareDatabase(t *testing.T) {
	if err := repo.Prepare(databasePath); err != nil {
		t.Error(err)
	}
}

type userTestData struct {
	nickname string
	password []byte
}

var testData = []userTestData{
	{nickname: "cactus", password: []byte("supercactushero")},
	{nickname: "gurenlagan", password: []byte("iveneverwatchedgurenlagan")},
	{nickname: "mikhail", password: []byte("123")},
}

func TestInsertUser(t *testing.T) {
	for _, data := range testData {
		_, err := repo.InsertUser(data.nickname, data.password)
		if err != nil {
			t.Errorf("%v, params: %v", err, data)
		}
	}
}

func TestSelectUser(t *testing.T) {
	for _, data := range testData {
		user, err := repo.SelectUser(data.nickname)
		if err != nil {
			t.Errorf("%v, params: %q", err, data.nickname)

		} else if !reflect.DeepEqual(user.Password, data.password) {
			t.Errorf("passwords are not equal: from db %v, test data %v", user.Password, data.password)
		}
	}
}

func TestCloseDatabase(t *testing.T) {
	if err := repo.Close(); err != nil {
		t.Error(err)
	}

	if err := os.Remove(path.FromRoot(databasePath)); err != nil {
		t.Error(err)
	}
}
