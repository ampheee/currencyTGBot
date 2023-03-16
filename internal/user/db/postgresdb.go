package db

import (
	storage "_entryTask/internal/user"
	"_entryTask/pkg/logger"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

type DB struct {
	pool   *pgxpool.Pool
	logger zerolog.Logger
}

func (db *DB) Create(ctx context.Context) (string, error) {
	//TODO Implement me!
	return "", nil
}

func (db *DB) FindOne(ctx context.Context) (string, error) {
	//TODO Implement me!
	return "", nil
}

func (db *DB) FindAll(ctx context.Context) (data []string, err error) {
	//TODO Implement me!
	return nil, nil
}

func (db *DB) FindByDate(ctx context.Context) (data []string, err error) {
	//TODO Implement me!
	return nil, nil
}

func NewStorage(pool *pgxpool.Pool) storage.Storage {
	return &DB{pool: pool, logger: logger.GetLogger()}
}
