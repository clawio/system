package main

import (
	"fmt"
	"net"
	"os"

	"github.com/clawio/system/pb"
	"github.com/clawio/system/pkg/config"
	"github.com/clawio/system/pkg/errors"
	"github.com/clawio/system/pkg/sec"
	"github.com/clawio/system/pkg/session"
	"github.com/clawio/system/pkg/user"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/levels"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"gopkg.in/ini.v1"
)

type server struct {
	logger         levels.Levels
	userManager    user.UserManager
	sessionManager session.SessionManager
}

func newServer(logger levels.Levels, um user.UserManager, sm session.SessionManager) *server {
	return &server{
		logger:         logger,
		userManager:    um,
		sessionManager: sm,
	}
}
func (s *server) Authenticate(ctx context.Context, creds *pb.Credentials) (*pb.Token, error) {
	user, err := s.userManager.Authenticate(creds)
	if err != nil {
		s.logger.Error().Log("msg", "error authenticating", "err", err)
		return nil, errors.NewError(errors.UnauthorizedErrorCode, "")
	}
	s.logger.Info().Log("msg", "authentication ok", "account_id", user.AccountId)
	token, err := s.sessionManager.GenerateSessionToken(user)
	if err != nil {
		s.logger.Error().Log("msg", "error generating token", "err", err)
		return nil, errors.NewError(errors.UnauthorizedErrorCode, "")
	}
	res := &pb.Token{Value: token}
	return res, nil
}

func (s *server) Whoami(ctx context.Context, _ *pb.Empty) (*pb.User, error) {
	user, err := sec.Authorize(ctx, s.sessionManager)
	if err != nil {
		s.logger.Error().Log("msg", "authentication failed", "err", err)
		return nil, err
	}
	return user, nil
}

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	leveledLogger := levels.New(logger)

	config, err := ini.LooseLoad("/etc/clawio/clawioauthserver.conf", "./clawioauthserver.conf", config.DefaultConfig)
	if err != nil {
		leveledLogger.Crit().Log("msg", "can't read config", "err", err)
		os.Exit(1)
	}

	um := user.GetUserManagerFromConfig(config)
	sm := session.GetSessionManagerFromConfig(config)
	server := newServer(leveledLogger, um, sm)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Section("").Key("server_port").MustInt(1502)))
	if err != nil {
		leveledLogger.Crit().Log("err", err)
		os.Exit(1)
	}

	leveledLogger.Info().Log("msg", "listening", "addr", config.Section("").Key("server_port").MustInt(1502))

	srv := grpc.NewServer()
	pb.RegisterAuthServer(srv, server)
	srv.Serve(lis)
}
