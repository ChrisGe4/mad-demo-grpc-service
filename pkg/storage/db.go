package storage

import (
	"context"

	"github.com/chrisge4/mad-demo-grpc-service/pkg/pb"
)

type Database interface {
	//Get(ctx context.Context, uri string) (io.ReadCloser, error)
	Add(context.Context, *pb.Task) error
	List(context.Context, *pb.Category) ([]*pb.Task, error)
	Close() error
}
