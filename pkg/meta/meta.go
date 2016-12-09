package meta

import (
	"context"
	"github.com/clawio/system/pb"
)

type MetaDataManager interface {
	Info(ctx context.Context, path string) (*pb.MetaData, error)
	ListContainer(ctx context.Context, path string) ([]*pb.MetaData, error)
	Remove(ctx context.Context, path string) error
	Move(ctx context.Context, from, to string) error
	CreateContainer(ctx context.Context, path string) error
}
