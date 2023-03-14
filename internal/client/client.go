package client

import (
	"_entryTask/internal/handlers"
	"context"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
)

func Run(ctx context.Context, tgbot *tg.Client) error {
	router := tgb.NewRouter()
	handlers.InitHandlers(router)
	err := tgb.NewPoller(router, tgbot).Run(ctx)
	return err
}
