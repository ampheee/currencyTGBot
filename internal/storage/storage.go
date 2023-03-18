package storage

import (
	"context"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
)

type Storage interface {
	Create(ctx context.Context, update *tgb.Update) error
	FindAllRequestsById(ctx context.Context, userId tg.UserID) (data []string, err error)
	FindFirstRequest(ctx context.Context, userId int) (string, error)
	AddNewRequest(ctx context.Context, update *tgb.Update) error
}
