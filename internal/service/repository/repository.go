package service

import (
	"_entryTask/config"
	"_entryTask/internal/coin"
	"_entryTask/internal/service"
	"_entryTask/internal/storage"
	botErrors "_entryTask/pkg/constants/errs"
	"_entryTask/pkg/logger"
	"context"
	"fmt"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/rs/zerolog"
	"strings"
)

type Service struct {
	storage   storage.Storage
	botLogger zerolog.Logger
	config    config.Config
}

func (s *Service) CreateUser(ctx context.Context) {

}

func (s *Service) GetStats(ctx context.Context, id int, msg string) {

}

func (s *Service) ClearStats(ctx context.Context, id int) bool {

	return true
}

func (s *Service) GetInfoAboutCurrency(ctx context.Context, update *tgb.Update) (currency *coin.Currency, err int) {
	msgSlice := strings.Split(update.Message.Text, " ")
	if len(msgSlice) > 2 {
		s.botLogger.Warn().Msg(fmt.Sprintf("[CurrencyInfo] Invalid input \"%s\" from [%d %s]",
			update.Message.Text, update.Message.From.ID, update.Chat().Username))
		return nil, botErrors.InvalidInput
	}
	currency = coin.GetCurrencyInfo(s.config, msgSlice[1])
	if currency == nil {
		s.botLogger.Warn().Msg(fmt.Sprintf("[CurrencyInfo] Currency not found \"%s\" from [%d %s]",
			msgSlice[1], update.Message.From.ID, update.Chat().Username))
		return nil, botErrors.CurrencyNotFound
	}
	s.storage.FindAllRequestsById(ctx, update.Message.From.ID)
	s.storage.AddNewRequest(ctx)
	return currency, 0
}

func (s *Service) ExchangeTwoCurrencies(ctx context.Context) {
	s.storage.AddNewRequest(ctx)
}

func NewService(storage storage.Storage, config config.Config) service.UserService {
	return &Service{storage, logger.GetLogger(), config}
}
