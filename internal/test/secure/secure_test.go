package secure_test

import (
	"testing"

	"github.com/AuthorDriu/namespool/internal/secure"
)

type passwordTestData struct {
	passwordString    string
	passwordEnctypted []byte
}

var passwordData = []passwordTestData{
	{passwordString: "hellosuka"},
	{passwordString: "uliuliu"},
	{passwordString: "notevenapassword"},
	{passwordString: "fourthpassword"},
	{passwordString: "ihaveavadfantasy"},
	{passwordString: "itsdifficulttocreateapasswordfortests"},
}

func TestPasswordEncrypting(t *testing.T) {
	for i := range passwordData {
		encrypted, err := secure.EnctyptPassord(passwordData[i].passwordString)
		if err != nil {
			t.Error(err)
		} else if len(encrypted) == 0 {
			t.Error("password enctypting returns zero length slice")
		}
		passwordData[i].passwordEnctypted = encrypted
	}
}

func TestValidatePassword(t *testing.T) {
	for i := range passwordData {
		if !secure.ValidatePassword(passwordData[i].passwordString, passwordData[i].passwordEnctypted) {
			t.Errorf("cannot validate password: %v", passwordData[i])
		}
	}
}

type tokenTestData struct {
	nickname string
	token    string
}

var tokenData = []tokenTestData{
	{nickname: "cactus"},
	{nickname: "voldemar"},
	{nickname: "yohan"},
	{nickname: "budimir"},
	{nickname: "sviatoslav"},
}

func TestGenerateToken(t *testing.T) {
	for i := range tokenData {
		token, err := secure.GenerateJWT(tokenData[i].nickname)
		if err != nil {
			t.Error(err)
		}
		tokenData[i].token = token
	}
}

func TestParseToken(t *testing.T) {
	for _, data := range tokenData {
		nickname, err := secure.ParseJWT(data.token)
		if err != nil {
			t.Error(err)
		} else if data.nickname != nickname {
			t.Errorf("wrong nickname! expected: %q, got: %q", data.nickname, nickname)
		}
	}
}
