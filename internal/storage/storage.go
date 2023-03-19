package storage

import (
	"context"
	"github.com/mr-linch/go-tg"
)

type Storage interface {
	CreateUser(ctx context.Context, user User) error
	FindAllRequestsById(ctx context.Context, userId tg.UserID) ([]string, error)
	DeleteUserStatsById(ctx context.Context, userId tg.UserID) error
	AddNewRequest(ctx context.Context, record RequestRecord) error
}
