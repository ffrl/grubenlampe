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

// ConnectTestDatabase returns a scope with a test database
func ConnectTestDatabase(t *testing.T) *TestDatabaseScope {
	file := uuid.New().String() + "-test.db"
	db, err := database.Connect("sqlite3", file, database.WithDebug())
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

	if err := s.addTestOrg(); err != nil {
		t.Fatalf("could not add fftest org. %s", err)
	}

	if err := s.addTestASN(); err != nil {
		t.Fatalf("could not add test ASN. %s", err)
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

func (s *TestDatabaseScope) addTestOrg() error {
	o := &database.Org{
		Name:      "Freifunk Test e.V.",
		ShortName: "fftest",
	}
	return s.db.Orgs().Save(o)
}

func (s *TestDatabaseScope) addTestASN() error {
	o, err := s.db.Orgs().GetByShortName("fftest")
	if err != nil {
		return err
	}

	a := &database.ASN{
		Org:     o,
		ASN:     201701,
		Checked: true,
	}
	return s.db.ASNs().Save(a)
}

func (s *TestDatabaseScope) Close() {
	s.db.Close()
	os.Remove(s.dbFile)
}

func (s *TestDatabaseScope) DB() *database.Connection {
	return s.db
}
