package data

import (
	"io"
	"context"
)

type DataManager interface {
	PutFile(ctx context.Context,  path string, r io.Reader, clientXS string) error
	GetFile(ctx context.Context, path string) (io.Reader, error)
}
