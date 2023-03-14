package postgresql

import (
	"_entryTask/config"
	"_entryTask/pkg/logger"
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

func InitDB(config config.Config, ctx context.Context) (db *sqlx.DB) {
	log := logger.GetLogger()
	err := ConnectWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, time.Second*3)
		defer cancel()
		var err error
		db, err = sqlx.ConnectContext(ctx, "pgx", config.DBURL)
		if err != nil {
			return err
		}
		return nil
	}, 5, time.Second*3)
	if err != nil {
		log.Fatal().Err(err).Msg("I couldn`t connect to DB. Is DBURL correct?")
	}
	return
}

func ConnectWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts < 0 {
		err = fn()
		if err != nil {
			attempts--
			time.Sleep(delay)
			continue
		}
		return nil
	}
	return
}
