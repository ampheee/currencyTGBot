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
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/rs/zerolog"
	"strings"
	"time"
)

type Service struct {
	storage   storage.Storage
	botLogger zerolog.Logger
	config    config.Config
}

func (s *Service) SendHelp(ctx context.Context, update *tgb.Update) string {
	str := tg.HTML.Text(
		tg.HTML.Bold("All existing commands:\n"),
		tg.HTML.Code(
			fmt.Sprintf("%s: %s\n%s: %s\n%s:%s\n%s: %s\n%s: %s",
				"\t/start", tg.HTML.Bold("calling greetings with credits"),
				"\t/stats", tg.HTML.Bold("shows your statistic for all time"),
				"\t/exchange", tg.HTML.Bold("exchanges one currency into another"),
				"\t/currencyInfo", tg.HTML.Bold("get entered currency info from coinAPI"),
				"\t/clearstats", tg.HTML.Bold("clear all your stats\n"),
			)),
		tg.HTML.Bold("Example usage:"),
		tg.HTML.Code("\t/exchange BTC USD\n\t/currencyinfo BTC"))
	err := s.storage.AddNewRequest(ctx, storage.RequestRecord{User: storage.User{
		Id:        update.Message.From.ID,
		Username:  update.Message.From.Username,
		FirstName: update.Message.From.FirstName,
		LastName:  update.Message.From.LastName,
	}, RequestType: "/help", RequestTime: time.Unix(int64(update.Message.Date), 0)})
	if err != nil {
		s.botLogger.Warn().Err(err).Msg("[HelpCommand] Unable to add request into db")
	}
	return str
}

func (s *Service) SendStart(ctx context.Context, update *tgb.Update) string {
	str := tg.HTML.Text(
		tg.HTML.Bold("\n 💸 I`m currency bot 💸\n"),
		tg.HTML.Italic("🚀 Powered by ampheee 🚀\n",
			tg.HTML.Underline(tg.HTML.Link("my github", "github.com/ampheee"))),
		tg.HTML.Line("It`s entry-task product"))
	err := s.storage.AddNewRequest(ctx, storage.RequestRecord{User: storage.User{
		Id:        update.Message.From.ID,
		Username:  update.Message.From.Username,
		FirstName: update.Message.From.FirstName,
		LastName:  update.Message.From.LastName,
	}, RequestType: "/start", RequestTime: time.Unix(int64(update.Message.Date), 0)})
	if err != nil {
		s.botLogger.Warn().Err(err).Msg("[StartCommand] Unable to add request into db")
	}
	return str
}

func (s *Service) GetStats(ctx context.Context, update *tgb.Update) []string {
	var request = storage.RequestRecord{
		User: storage.User{
			Id:        update.Message.From.ID,
			Username:  update.Message.From.Username,
			FirstName: update.Message.From.FirstName,
			LastName:  update.Message.From.LastName,
		},
		RequestType: update.Message.Text,
		RequestTime: time.Unix(int64(update.Message.Date), 0),
	}
	data, err := s.storage.FindAllRequestsById(ctx, request)
	if err != nil {
		s.botLogger.Warn().Err(err).Msg("[serviceGetStats] Can`t get stats")
	}
	return data
}

func (s *Service) ClearStats(ctx context.Context, userId tg.UserID) error {
	return s.storage.DeleteUserStatsById(ctx, userId)
}

func (s *Service) GetInfoAboutCurrency(ctx context.Context, update *tgb.Update) (currency *coin.Currency, err int) {
	msgSlice := strings.Split(update.Message.Text, " ")
	if len(msgSlice) != 2 {
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
	var requestRecord = storage.RequestRecord{
		User: storage.User{
			Id:        update.Message.From.ID,
			Username:  update.Message.From.Username,
			FirstName: update.Message.From.FirstName,
			LastName:  update.Message.From.LastName,
		},
		RequestType: msgSlice[0],
		RequestTime: time.Unix(int64(update.Message.Date), 0),
		RequestArgs: msgSlice[1],
	}
	dbErr := s.storage.AddNewRequest(ctx, requestRecord)
	if dbErr != nil {
		s.botLogger.Warn().Err(dbErr)
	}
	return currency, 0
}

func NewService(storage storage.Storage, config config.Config) service.UserService {
	return &Service{storage, logger.GetLogger(), config}
}
