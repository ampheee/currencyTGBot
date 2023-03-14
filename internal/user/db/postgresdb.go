package db

import (
	storage "_entryTask/internal/user"
	"_entryTask/pkg/logger"
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type DB struct {
	database *sqlx.DB
	logger   zerolog.Logger
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

func NewStorage(database *sqlx.DB) storage.Storage {
	return &DB{database: database, logger: logger.GetLogger()}
}
