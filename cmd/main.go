package main

import (
	"_entryTask/config"
	"_entryTask/internal/client"
	service "_entryTask/internal/service/repository"
	"_entryTask/internal/storage/repository"
	"_entryTask/pkg/logger"
	"_entryTask/pkg/postgresql"
	"context"
	"flag"
	"github.com/mr-linch/go-tg"
)

var skipMigrations = flag.Bool("migration", false, "set value true, if migration needed")

func main() {
	flag.Parse()
	ctx := context.Background()
	v := config.LoadConfig()
	config := config.ParseConfig(v)
	botLogger := logger.GetLogger()
	pool := postgresql.GetPool(*config, ctx)
	if *skipMigrations != false {
		postgresql.MigrateDatabase(ctx, pool)
	}
	userStorage := repository.NewStorage(pool)
	userService := service.NewService(userStorage, *config)
	tgBot := tg.New(*config.Tokens.TgToken)
	botLogger.Info().Msg("Bot was started!")
	botLogger.Fatal().Err(client.Run(ctx, tgBot, userService))
}
