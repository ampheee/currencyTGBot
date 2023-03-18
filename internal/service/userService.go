package service

import (
	"_entryTask/internal/coin"
	"context"
	"github.com/mr-linch/go-tg/tgb"
)

type UserService interface {
	CreateUser(ctx context.Context)
	GetInfoAboutCurrency(ctx context.Context, update *tgb.Update) (currency *coin.Currency, err int)
	ExchangeTwoCurrencies(ctx context.Context)
	GetStats(ctx context.Context, id int, msg string)
	ClearStats(ctx context.Context, id int) bool
}
