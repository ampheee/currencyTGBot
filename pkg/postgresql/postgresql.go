package postgresql

import (
	"_entryTask/config"
	"_entryTask/pkg/logger"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/tern/migrate"
	"time"
)

func GetPool(config config.Config, ctx context.Context) (pool *pgxpool.Pool) {
	botLogger := logger.GetLogger()
	err := ConnectWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, time.Second*3)
		defer cancel()
		var err error
		pool, err = pgxpool.Connect(ctx, config.DBURL)
		return err
	}, 5, time.Second*3)
	if err != nil {
		botLogger.Fatal().Err(err).Msg("I couldn`t connect to DB. Is DBURL correct?")
	}
	botLogger.Info().Msg(fmt.Sprintf("Pool connected"))
	return pool
}

func ConnectWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			attempts--
			time.Sleep(delay)
			continue
		}
		return nil
	}
	return
}

func MigrateDatabase(ctx context.Context, pool *pgxpool.Pool) {
	botLogger := logger.GetLogger()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		botLogger.Fatal().Err(err).Msg("Unable to acquire db connection")
	}
	migrator, err := migrate.NewMigrator(ctx, conn.Conn(), "version_01")
	if err != nil {
		botLogger.Fatal().Err(err).Msg("Unable to create migration")
	}
	err = migrator.LoadMigrations("./migrations")
	if err != nil {
		botLogger.Fatal().Err(err).Msg("Unable to load migrations")
	}
	err = migrator.Migrate(ctx)
	if err != nil {
		botLogger.Fatal().Err(err).Msg("Unable to migrate")
	}
	ver, err := migrator.GetCurrentVersion(ctx)
	if err != nil {
		botLogger.Fatal().Err(err).Msg("Unable to get current schema version")
	}
	botLogger.Info().Msg(fmt.Sprintf("Migration Done. Current schema version: %v", ver))
	defer conn.Release()
}
