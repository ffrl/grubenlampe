package server

import (
	"context"
	"testing"

	pb "github.com/ffrl/grubenlampe/api"
	"github.com/ffrl/grubenlampe/database"
	"github.com/ffrl/grubenlampe/testhelper"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	testCases := []struct {
		name  string
		email string
		f     func(*pb.GenericResponse, *database.Connection)
	}{
		{name: "should add user", email: "hans@wurst.de", f: func(res *pb.GenericResponse, db *database.Connection) {
			assertAddedUserExists(t, db)
		}},
		{name: "should not return success on existing user", email: "otto@ffrl.de", f: func(res *pb.GenericResponse, db *database.Connection) {
			assertAddUserFails(t, res, "User already exists")
		}},
		{name: "should return success", email: "hans@wurst.de", f: func(res *pb.GenericResponse, db *database.Connection) {
			assertAddUserReturnsSuccess(t, res)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(te *testing.T) {
			scope := testhelper.ConnectTestDatabase(t)
			defer scope.Close()

			db := scope.DB()
			s := &Server{db}

			req := &pb.AddUserRequest{
				Email:      tc.email,
				Password:   "passwort",
				RipeHandle: "xxx",
			}

			res, err := s.AddUser(context.TODO(), req)
			if err != nil {
				t.Fatalf("error while processing. %s", err)
			}

			tc.f(res, db)
		})
	}
}

func assertAddedUserExists(t *testing.T, db *database.Connection) {
	u, err := db.Users().GetByEmail("hans@wurst.de")
	if err != nil {
		t.Fatalf("could not get user added to database. %s", err)
	}

	assert.NotNil(t, u, "user should not be nil")
	assert.EqualValues(t, "hans@wurst.de", u.Email, "email")
	assert.EqualValues(t, "passwort", u.Password, "password")
	assert.EqualValues(t, "xxx", u.RIPEHandle, "ripe-handle")
}

func assertAddUserReturnsSuccess(t *testing.T, res *pb.GenericResponse) {
	assert.True(t, res.Success, "add user should succeed")
}

func assertAddUserFails(t *testing.T, res *pb.GenericResponse, message string) {
	assert.False(t, res.Success, "add user should not succeed")
	assert.EqualValues(t, message, res.Message)
}
