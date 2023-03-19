package postgresql

import (
	"_entryTask/config"
	"_entryTask/pkg/logger"
	"_entryTask/pkg/middleware"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/tern/migrate"
	"time"
)

func GetPool(config config.Config, ctx context.Context) (pool *pgxpool.Pool) {
	botLogger := logger.GetLogger()
	err := middleware.ConnectWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, time.Second*3)
		defer cancel()
		var err error
		pool, err = pgxpool.Connect(ctx, fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s",
			*config.Postgres.User,
			*config.Postgres.Password,
			*config.Postgres.Host,
			*config.Postgres.Port,
			*config.Postgres.DbName,
			"sslmode=disable&pool_max_conns=10"))
		return err
	}, 3, time.Second*3)
	if err != nil {
		botLogger.Fatal().Err(err).Msg("Unable connect to DB")
	}
	botLogger.Info().Msg(fmt.Sprintf("Pool connected"))
	return pool
}

func MigrateDatabase(ctx context.Context, pool *pgxpool.Pool) {
	botLogger := logger.GetLogger()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		botLogger.Fatal().Err(err).Msg("Unable to acquire repository connection")
	}
	migrator, err := migrate.NewMigrator(ctx, conn.Conn(), "public.version")
	if err != nil {
		botLogger.Fatal().Err(err).Msg("Unable to create migration")
	}
	err = migrator.LoadMigrations("../pkg/postgresql/migrations")
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
	conn.Release()
}
