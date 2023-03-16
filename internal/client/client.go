package client

import (
	"_entryTask/internal/handlers"
	"_entryTask/internal/service"
	"_entryTask/pkg/logger"
	"context"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
)

func Run(ctx context.Context, tgBot *tg.Client, service service.UserService) error {
	botLogger := logger.GetLogger()
	router := tgb.NewRouter()
	handlers.InitHandlers(ctx, router, service)
	_, err := tgBot.GetMe().Do(ctx)
	if err != nil {
		botLogger.Fatal().Err(err).Msg("Unable to start bot. Check token and start again")
	}
	err = tgb.NewPoller(router, tgBot).Run(ctx)
	return err
}
