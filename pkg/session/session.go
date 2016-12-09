package session

import (
	"github.com/clawio/system/pb"
	"gopkg.in/ini.v1"
)

type SessionManager interface {
	// CreateSessionToken creates a temporary ticket for the user
	GenerateSessionToken(user *pb.User) (string, error)

	// DecodeSessionToken decodes the ticket into a user
	DecodeSessionToken(ticket string) (*pb.User, error)
}

func GetSessionManagerFromConfig(config *ini.File) SessionManager {
	return NewJWTSessionManager(config)
}
