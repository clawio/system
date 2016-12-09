package sec

import (
	"golang.org/x/net/context"
	"github.com/clawio/system/pb"
	"google.golang.org/grpc/metadata"
	"github.com/clawio/system/pkg/session"
	"github.com/clawio/system/pkg/errors"
)

func Authorize(ctx context.Context, sessionManager session.SessionManager) (*pb.User, error) {
	if md, ok := metadata.FromContext(ctx); ok {
		if len(md["token"]) > 0 && md["token"][0] != "" {
			user, err := sessionManager.DecodeSessionToken(md["token"][0])
			if err != nil {
				return nil, errors.NewError(errors.UnauthorizedErrorCode, "token is not valid")
			}
			return user, nil
		}
		return nil, errors.NewError(errors.UnauthorizedErrorCode, "no token provided")
	}
	return nil, errors.NewError(errors.UnauthorizedErrorCode, "no client metadata present")
}
