package testhelper

import (
	"os"
	"testing"

	"github.com/ffrl/grubenlampe/database"
	"github.com/google/uuid"
)

type TestDatabaseScope struct {
	dbFile string
	db     *database.Connection
	t      *testing.T
}

func ConnectTestDatabase(t *testing.T) *TestDatabaseScope {
	file := uuid.New().String() + "-test.db"
	db, err := database.Connect("sqlite3", file)
	if err != nil {
		t.Fatalf("could not connect to test database. %s", err)
	}

	s := &TestDatabaseScope{
		db:     db,
		dbFile: file,
		t:      t,
	}

	if err := s.addTestUserNormal(); err != nil {
		t.Fatalf("could not add normal user. %s", err)
	}

	if err := s.addTestUserAdmin(); err != nil {
		t.Fatalf("could not add admin user. %s", err)
	}

	return s
}

func (s *TestDatabaseScope) addTestUserNormal() error {
	u := &database.User{
		Email:      "otto@ffrl.de",
		Password:   "normal",
		RIPEHandle: "OTTO-RIPE",
	}
	return s.db.Users().Save(u)
}

func (s *TestDatabaseScope) addTestUserAdmin() error {
	u := &database.User{
		Email:      "admin@ffrl.de",
		Password:   "admin",
		RIPEHandle: "admin1234",
	}
	return s.db.Users().Save(u)
}

func (s *TestDatabaseScope) Close() {
	s.db.Close()
	os.Remove(s.dbFile)
}

func (s *TestDatabaseScope) DB() *database.Connection {
	return s.db
}
