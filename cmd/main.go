package main

import (
	"_entryTask/config"
	"_entryTask/internal/client"
	"_entryTask/pkg/logger"
	"_entryTask/pkg/postgresql"
	"context"
	"flag"
	"github.com/mr-linch/go-tg"
)

var skipMigrations = flag.Bool("migration", false, "set value true, if migration needed")

func main() {
	flag.Parse()
	Config := config.Load()
	ctx := context.Background()
	botLogger := logger.GetLogger()
	pool := postgresql.GetPool(Config, ctx)
	if *skipMigrations != false {
		postgresql.MigrateDatabase(ctx, pool)
	}
	tgBot := tg.New(Config.TGBotToken)
	//strge := db.NewStorage(database)
	botLogger.Fatal().Err(client.Run(ctx, tgBot))
}
