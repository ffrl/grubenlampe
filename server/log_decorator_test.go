package server

import (
	"context"
	"testing"

	pb "github.com/ffrl/grubenlampe/api"
	"github.com/ffrl/grubenlampe/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestWriteLog(t *testing.T) {
	scope := testhelper.ConnectTestDatabase(t)
	defer scope.Close()

	db := scope.DB()
	s := WithLogging(db, &Server{db})

	req := &pb.AddUserRequest{
		Email:      "log@ffrl.de",
		Password:   "xxx",
		RipeHandle: "log-123",
	}

	_, err := s.AddUser(context.TODO(), req)
	if err != nil {
		t.Fatal(err)
	}

	l, err := db.Logs().GetLog()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(l), "log message count")
}
