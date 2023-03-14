package storage

import (
	"context"
	"time"
)

type Storage interface {
	Create(ctx context.Context) (string, error)
	FindAll(ctx context.Context) (data []string, err error)
	FindOne(ctx context.Context) (string, error)
	FindByDate(ctx context.Context) (data []string, err error)
}

type Record struct {
	UserName       string
	Command        string
	FirstCurrency  string
	SecondCurrency string
	Date           time.Time
	Result         string
}
