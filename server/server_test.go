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
		name      string
		email     string
		assertion func(*pb.GenericResponse, *database.Connection)
	}{
		{name: "should add user", email: "hans@wurst.de", assertion: func(res *pb.GenericResponse, db *database.Connection) {
			assertAddedUserExists(t, db)
		}},
		{name: "should not return success on existing user", email: "otto@ffrl.de", assertion: func(res *pb.GenericResponse, db *database.Connection) {
			assertFails(t, res, "User already exists")
		}},
		{name: "should return success", email: "hans@wurst.de", assertion: func(res *pb.GenericResponse, db *database.Connection) {
			assertReturnsSuccess(t, res)
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

			tc.assertion(res, db)
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

func assertReturnsSuccess(t *testing.T, res *pb.GenericResponse) {
	assert.True(t, res.Success, "add user should succeed")
}

func assertFails(t *testing.T, res *pb.GenericResponse, message string) {
	assert.False(t, res.Success, "add user should not succeed")
	assert.EqualValues(t, message, res.Message)
}

func TestAddOrg(t *testing.T) {
	testCases := []struct {
		name      string
		shortName string
		assertion func(*pb.GenericResponse, *database.Connection)
	}{
		{name: "should add org", shortName: "ffneu", assertion: func(res *pb.GenericResponse, db *database.Connection) {
			assertAddedOrgExists(t, db)
		}},
		{name: "should return success", shortName: "ffneu", assertion: func(res *pb.GenericResponse, db *database.Connection) {
			assertReturnsSuccess(t, res)
		}},
		{name: "should fail if org short name exists", shortName: "fftest", assertion: func(res *pb.GenericResponse, db *database.Connection) {
			assertFails(t, res, "Org already exists")
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(te *testing.T) {
			scope := testhelper.ConnectTestDatabase(t)
			defer scope.Close()

			db := scope.DB()
			s := &Server{db}

			req := &pb.AddOrgRequest{
				Name:      "Freifunk Test",
				ShortName: tc.shortName,
			}

			res, err := s.AddOrg(context.TODO(), req)
			if err != nil {
				t.Fatalf("error while processing. %s", err)
			}

			tc.assertion(res, db)
		})
	}
}

func assertAddedOrgExists(t *testing.T, db *database.Connection) {
	o, err := db.Orgs().GetByShortName("ffneu")
	if err != nil {
		t.Fatalf("could not get org added to database. %s", err)
	}

	assert.NotNil(t, o, "org should not be nil")
	assert.EqualValues(t, "Freifunk Test", o.Name, "name")
	assert.EqualValues(t, "ffneu", o.ShortName, "short-name")
}
