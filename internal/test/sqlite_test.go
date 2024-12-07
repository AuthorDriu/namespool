package sqlite_test

import (
	"os"
	"reflect"
	"testing"

	dbTypes "github.com/AuthorDriu/namespool/internal/repository"
	db "github.com/AuthorDriu/namespool/internal/repository/sqlite"
	"github.com/AuthorDriu/namespool/pkg/path"
)

const databasePath string = "store/test_database.db"

func TestPrepareDatabase(t *testing.T) {
	if err := db.Prepare(databasePath); err != nil {
		t.Error(err)
	}
}

type userTestData struct {
	nickname string
	password []byte
}

var userData = []userTestData{
	{nickname: "cactus", password: []byte("supercactushero")},
	{nickname: "gurenlagan", password: []byte("iveneverwatchedgurenlagan")},
	{nickname: "mikhail", password: []byte("123")},
	{nickname: "егорычзмей", password: []byte("esnake")},
}

func TestInsertUser(t *testing.T) {
	for _, data := range userData {
		_, err := db.InsertUser(data.nickname, data.password)
		if err != nil {
			t.Errorf("%v, params: %v", err, data)
		}
	}
}

func TestSelectUser(t *testing.T) {
	for _, data := range userData {
		user, err := db.SelectUser(data.nickname)
		if err != nil {
			t.Errorf("%v, params: %q", err, data.nickname)
		} else if !reflect.DeepEqual(user.Password, data.password) {
			t.Errorf("passwords are not equal: from db %v, test data %v", user.Password, data.password)
		}
	}
}

type ideaTestData struct {
	title       string
	description string
	access      dbTypes.AccessModifier
	owner       string
}

var ideaData = []ideaTestData{
	{title: "hug", description: "just hugs", access: dbTypes.IdeaPublic, owner: "cactus"},
	{title: "suka", description: "insulting project", access: dbTypes.IdeaPublic, owner: "cactus"},
	{title: "bliat", description: "second insulting project", access: dbTypes.IdeaPublic, owner: "cactus"},
	{title: "haha", description: "", access: dbTypes.IdeaPrivate, owner: "cactus"},
}

func TestInsertIdea(t *testing.T) {
	for _, idea := range ideaData {
		_, err := db.InsertIdea(idea.title, idea.description, idea.access, idea.owner)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestInsertNotUniqueIdea(t *testing.T) {
	idea := ideaData[0]
	_, err := db.InsertIdea(idea.title, idea.description, idea.access, idea.owner)
	if err == nil {
		t.Error("Inserted not unique idea")
	}
}

func TestSelectIdea(t *testing.T) {
	ideaData := ideaData[0]
	idea, err := db.SelectIdea(ideaData.owner, ideaData.title)
	if err != nil {
		t.Errorf("%v, params: {%q, %q}", err, ideaData.owner, ideaData.title)
	} else if idea.Title != ideaData.title {
		t.Errorf("Wrong result! %q != %q", idea.Title, ideaData.title)
	}
}

func TestSelectIdeasByUser(t *testing.T) {
	ideas, err := db.SelectIdeasByUser(ideaData[0].owner)
	if err != nil {
		t.Errorf("%v, params: %q", err, ideaData[0].owner)
	} else if len(ideas) != len(ideaData) {
		t.Errorf("Wrong result (wrong length) %d != %d:", len(ideas), len(ideaData))
	}
}

func TestSelectPublicIdeasByUser(t *testing.T) {
	ideas, err := db.SelectPublicIdeasByUser(ideaData[0].owner)
	if err != nil {
		t.Errorf("%v, params: %q", err, ideaData[0].owner)
	} else if len(ideas) != len(ideaData)-1 {
		t.Errorf("Wrong result (wrong length) %d != %d:", len(ideas), len(ideaData)-1)
	}
}

func TestDeleteIdea(t *testing.T) {
	for _, idea := range ideaData {
		err := db.DeleteIdea(idea.owner, idea.title)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestCloseDatabase(t *testing.T) {
	if err := db.Close(); err != nil {
		t.Error(err)
	}

	if err := os.Remove(path.FromRoot(databasePath)); err != nil {
		t.Error(err)
	}
}
