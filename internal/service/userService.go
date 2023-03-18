package service

import (
	"_entryTask/internal/coin"
	"context"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
)

type UserService interface {
	GetInfoAboutCurrency(ctx context.Context, update *tgb.Update) (currency *coin.Currency, err int)
	GetStats(ctx context.Context, update *tgb.Update) []string
	ClearStats(ctx context.Context, userId tg.UserID) error
	SendHelp(ctx context.Context, update *tgb.Update) string
	SendStart(ctx context.Context, update *tgb.Update) string
}
