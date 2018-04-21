package server

import (
	"context"
	"testing"

	pb "github.com/ffrl/grubenlampe/api"
	"github.com/ffrl/grubenlampe/testhelper"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	scope := testhelper.ConnectTestDatabase(t)
	defer scope.Close()

	db := scope.DB()
	s := &Server{db}

	req := &pb.AddUserRequest{
		Email:      "hans@wurst.de",
		Password:   "passwort",
		RipeHandle: "xxx",
	}

	_, err := s.AddUser(context.TODO(), req)
	if err != nil {
		t.Fatalf("error while processing. %s", err)
	}

	u, err := db.Users().GetByEmail("hans@wurst.de")
	if err != nil {
		t.Fatalf("could not get user added to database. %s", err)
	}

	assert.NotNil(t, u, "user should not be nil")
	assert.EqualValues(t, "hans@wurst.de", u.Email, "email")
	assert.EqualValues(t, "passwort", u.Password, "password")
	assert.EqualValues(t, "xxx", u.RIPEHandle, "ripe-handle")
}
