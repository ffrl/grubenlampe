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

func assertReturnsSuccess(t *testing.T, res *pb.GenericResponse, err error) {
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, res.Success, "add user should succeed")
}

func assertFails(t *testing.T, res *pb.GenericResponse, message string, err error) {
	if err != nil {
		t.Fatal(err)
	}

	assert.False(t, res.Success, "add user should not succeed")
	assert.EqualValues(t, message, res.Message)
}

func assertErrorNotNil(t *testing.T, err error) {
	assert.NotNil(t, err, "error expected")
}

func TestAddUser(t *testing.T) {
	testCases := []struct {
		name      string
		email     string
		assertion func(*pb.GenericResponse, *database.Connection, error, *testing.T)
	}{
		{
			name:  "should add user",
			email: "hans@wurst.de",
			assertion: func(res *pb.GenericResponse, db *database.Connection, err error, te *testing.T) {
				assertAddedUserExists(t, db, err)
			},
		},
		{
			name:  "should not return success on existing user",
			email: "otto@ffrl.de",
			assertion: func(res *pb.GenericResponse, db *database.Connection, err error, te *testing.T) {
				assertFails(t, res, "User already exists", err)
			},
		},
		{
			name:  "should return success",
			email: "hans@wurst.de",
			assertion: func(res *pb.GenericResponse, db *database.Connection, err error, te *testing.T) {
				assertReturnsSuccess(t, res, err)
			},
		},
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
			tc.assertion(res, db, err, te)
		})
	}
}

func assertAddedUserExists(t *testing.T, db *database.Connection, err error) {
	if err != nil {
		t.Fatal(err)
	}

	u, err := db.Users().GetByEmail("hans@wurst.de")
	if err != nil {
		t.Fatalf("could not get user added to database. %s", err)
	}

	assert.NotNil(t, u, "user should not be nil")
	assert.EqualValues(t, "hans@wurst.de", u.Email, "email")
	assert.EqualValues(t, "passwort", u.Password, "password")
	assert.EqualValues(t, "xxx", u.RIPEHandle, "ripe_handle")
}

func TestAddOrg(t *testing.T) {
	testCases := []struct {
		name      string
		shortName string
		assertion func(*pb.GenericResponse, *database.Connection, error, *testing.T)
	}{
		{
			name:      "should add org",
			shortName: "ffneu",
			assertion: func(res *pb.GenericResponse, db *database.Connection, err error, te *testing.T) {
				assertAddedOrgExists(te, db)
			},
		},
		{
			name:      "should return success",
			shortName: "ffneu",
			assertion: func(res *pb.GenericResponse, db *database.Connection, err error, te *testing.T) {
				assertReturnsSuccess(te, res, err)
			},
		},
		{
			name:      "should fail if org short name exists",
			shortName: "fftest",
			assertion: func(res *pb.GenericResponse, db *database.Connection, err error, te *testing.T) {
				assertFails(te, res, "Org already exists", err)
			},
		},
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

			tc.assertion(res, db, err, te)
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
	assert.EqualValues(t, "ffneu", o.ShortName, "short_name")
}

func TestAddASN(t *testing.T) {
	testCases := []struct {
		name      string
		asn       uint32
		org       string
		assertion func(*pb.GenericResponse, *database.Connection, error, *testing.T)
	}{
		{
			name: "should add ASN",
			asn:  12345,
			org:  "fftest",
			assertion: func(res *pb.GenericResponse, db *database.Connection, err error, te *testing.T) {
				assertAddedASNExists(te, db)
			},
		},
		{
			name: "should return success",
			asn:  12345,
			org:  "fftest",
			assertion: func(res *pb.GenericResponse, db *database.Connection, err error, te *testing.T) {
				assertReturnsSuccess(te, res, err)
			},
		},
		{
			name: "should fail if checked asn exists",
			asn:  201701,
			org:  "fftest",
			assertion: func(res *pb.GenericResponse, db *database.Connection, err error, te *testing.T) {
				assertFails(te, res, "ASN already exists", err)
			},
		},
		{
			name: "should fail if org does not exist",
			asn:  12345,
			org:  "ffuntest",
			assertion: func(res *pb.GenericResponse, db *database.Connection, err error, te *testing.T) {
				assertErrorNotNil(te, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(te *testing.T) {
			scope := testhelper.ConnectTestDatabase(t)
			defer scope.Close()

			db := scope.DB()
			s := &Server{db}

			req := &pb.AddASNRequest{
				Asn:          tc.asn,
				OrgShortName: tc.org,
			}

			res, err := s.AddASN(context.TODO(), req)
			tc.assertion(res, db, err, te)
		})
	}
}

func assertAddedASNExists(t *testing.T, db *database.Connection) {
	a, err := db.ASNs().GetByNumber(12345)
	if err != nil {
		t.Fatalf("could not get ASN added to database. %s", err)
	}

	assert.NotNil(t, a, "ASN should not be nil")
	assert.NotNil(t, a.Org, "Org should be not nil")
	assert.EqualValues(t, a.Org.ID, 1, "org_id")
	assert.Nil(t, a.CheckedBy, "checked_by")
}
