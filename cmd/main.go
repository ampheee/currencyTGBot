package main

import (
	"_entryTask/config"
	"_entryTask/internal/client"
	"_entryTask/pkg/logger"
	"context"
	"github.com/mr-linch/go-tg"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	Config := config.Load()
	ctx := context.Background()
	botLogger := logger.GetLogger()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill, syscall.SIGTERM)
	defer cancel()
	//database := postgresql.InitDB(Config, ctx)
	tgBot := tg.New(Config.TGBotToken)
	//strge := db.NewStorage(database)
	botLogger.Fatal().Err(client.Run(ctx, tgBot))
}
