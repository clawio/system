package user

import (
	"github.com/clawio/system/pb"
	"gopkg.in/ini.v1"
)

type UserManager interface {
	// Authenticate checks if the supplied credentials are
	// valid and returns the userID.
	Authenticate(credentials *pb.Credentials) (*pb.User, error)

	// GetUser returns information about an user
	GetUser(accountID string) (*pb.User, error)

	// GetUsers returns a list of users that match the
	// supplied filter
	GetUsers(filter string) ([]*pb.User, error)

	// UserExists returns true if the supplied username exists
	UserExists(accountID string) (bool, error)

	// GetNumberOfUsers returns the number of users in the backend
	GetNumberOfUsers(filter string) (int, error)
}

func GetUserManagerFromConfig(config *ini.File) UserManager {
	return NewMemoryUserManager(config)
}
