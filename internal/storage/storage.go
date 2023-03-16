package storage

import (
	"context"
)

type Storage interface {
	Create(ctx context.Context) (string, error)
	FindAll(ctx context.Context) (data []string, err error)
	FindOne(ctx context.Context) (string, error)
	AddNewRequest(ctx context.Context)
}
