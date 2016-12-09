package session

import (
	"fmt"
	"github.com/clawio/system/pb"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/ini.v1"
	"time"
)

type jwtManager struct {
	config *ini.File
}

func NewJWTSessionManager(config *ini.File) SessionManager {
	return &jwtManager{config}
}

func (u *jwtManager) GenerateSessionToken(user *pb.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account_id":   user.AccountId,
		"display_name": user.DisplayName,
		"exp":          time.Now().Unix() + 3600, // 1 hour
		"nbf":          time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(u.config.Section("").Key("session_manager_jwt_secret").MustString("")))
	return tokenString, err
}

func (u *jwtManager) DecodeSessionToken(ticket string) (*pb.User, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(ticket, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(u.config.Section("").Key("session_manager_jwt_secret").MustString("")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		u := &pb.User{
			AccountId:   claims["account_id"].(string),
			DisplayName: claims["display_name"].(string),
		}
		return u, nil
	}
	return nil, err

}
