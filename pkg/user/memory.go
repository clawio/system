package user

import (
	"errors"
	"strings"

	"github.com/clawio/system/pb"
	"gopkg.in/ini.v1"
)

func NewMemoryUserManager(config *ini.File) UserManager {
	users := parseUsers(config)
	return &memoryManager{conf: config, users: users}
}

type user struct {
	accountID   string
	displayName string
	password    string
}

func (u *user) toProto() *pb.User {
	return &pb.User{AccountId: u.accountID, DisplayName: u.displayName}
}

type memoryManager struct {
	conf  *ini.File
	users []*user
}

func (u *memoryManager) GetNumberOfUsers(filter string) (int, error) {
	return len(u.users), nil
}

func (u *memoryManager) UserExists(accountID string) (bool, error) {
	for _, user := range u.users {
		if user.accountID == accountID {
			return true, nil
		}
	}
	return false, nil
}

func (u *memoryManager) GetUser(accountID string) (*pb.User, error) {
	for _, user := range u.users {
		if user.accountID == accountID {
			return user.toProto(), nil
		}
	}
	return nil, errors.New("user not found")
}

func (u *memoryManager) GetUsers(filter string) ([]*pb.User, error) {
	users := []*pb.User{}
	for _, user := range u.users {
		if strings.Contains(user.accountID, filter) || strings.Contains(user.displayName, filter) {
			users = append(users, user.toProto())
		}
	}
	return users, nil
}

func (u *memoryManager) Authenticate(credentials *pb.Credentials) (*pb.User, error) {
	accountID, pwd := u.parseCreds(credentials.Value)
	for _, user := range u.users {
		if user.accountID == accountID && user.password == pwd {
			return u.GetUser(user.accountID)
		}
	}
	return nil, errors.New("accountID/password don't match")
}

func (u *memoryManager) parseCreds(creds string) (string, string) {
	els := strings.Split(creds, ":")
	if len(els) == 0 {
		return "", ""
	} else if len(els) == 1 {
		return els[0], ""
	} else {
		return els[0], els[1]
	}
}

// parseUsers parses user.memory.users
func parseUsers(config *ini.File) []*user {
	raw := config.Section("").Key("user_manager_memory_users").MustString("")
	els := strings.Split(raw, ",")
	users := []*user{}
	for _, u := range els {
		sels := strings.Split(u, ":")
		if len(sels) == 3 {
			if sels[0] != "" && sels[1] != "" && sels[2] != "" {
				users = append(users, &user{accountID: sels[0], password: sels[1], displayName: sels[2]})
			}
		}
	}
	return users
}
