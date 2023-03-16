package service

import (
	"_entryTask/internal/coin"
	"context"
	"github.com/mr-linch/go-tg/tgb"
)

type UserService interface {
	CreateUser()
	GetInfoAboutCurrency(ctx context.Context, update *tgb.Update) (currency *coin.Currency, err int)
	ExchangeTwoCurrencies()
	GetStats(id int, msg string)
	ClearStats(id int) bool
}
