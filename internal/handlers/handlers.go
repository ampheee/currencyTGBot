package handlers

import (
	"_entryTask/internal/service"
	botErrors "_entryTask/pkg/constants/errs"
	"_entryTask/pkg/logger"
	"context"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
)

func InitHandlers(ctx context.Context, router *tgb.Router, service service.UserService) {
	StartCommand(router, service)
	HelpCommand(router, service)
	StatsCommand(router, service)
	CurrencyInfoCommand(router, service)
	ClearStatsCommand(router, service)
	UnknownCommand(router)
}

func StartCommand(router *tgb.Router, userService service.UserService) {
	botLogger := logger.GetLogger()
	router.Message(func(ctx context.Context, update *tgb.MessageUpdate) error {
		botLogger.Info().Msg(fmt.Sprintf("[Start] Fetched \"%s\" from [%d %v]",
			update.Update.Message.Text,
			update.Message.From.ID,
			update.Message.From.Username))
		err := update.Answer(userService.SendStart(ctx, update.Update)).ParseMode(tg.HTML).DoVoid(ctx)
		return err
	}, tgb.Command("start", tgb.WithCommandIgnoreCase(true)))
}

func HelpCommand(router *tgb.Router, service service.UserService) {
	botLogger := logger.GetLogger()
	router.Message(func(ctx context.Context, update *tgb.MessageUpdate) error {
		botLogger.Info().Msg(fmt.Sprintf("[Start] Fetched \"%s\" from [%d %v]",
			update.Update.Message.Text,
			update.Message.From.ID,
			update.Message.From.Username))
		err := update.Answer(
			service.SendHelp(ctx, update.Update),
		).ParseMode(tg.HTML).DoVoid(ctx)
		return err
	}, tgb.Command("help", tgb.WithCommandIgnoreCase(true)))
}

func StatsCommand(router *tgb.Router, service service.UserService) {
	botLogger := logger.GetLogger()
	router.Message(func(ctx context.Context, update *tgb.MessageUpdate) error {
		botLogger.Info().Msg(fmt.Sprintf("[Stats] Fetched \"%s\" from [%d %v]",
			update.Update.Message.Text,
			update.Message.From.ID,
			update.Message.From.Username))
		data := service.GetStats(ctx, update.Update)
		err := update.Answer(
			tg.HTML.Bold(fmt.Sprintf(" Total requests: %d\n", len(data)), "Your commands:\n",
				tg.HTML.Code(service.GetStats(ctx, update.Update)...))).
			ParseMode(tg.HTML).DoVoid(ctx)
		return err
	}, tgb.Command("stats"))
}

func CurrencyInfoCommand(router *tgb.Router, userService service.UserService) {
	botLogger := logger.GetLogger()
	router.Message(func(ctx context.Context, update *tgb.MessageUpdate) error {
		botLogger.Info().Msg(fmt.Sprintf("[Exchange] Fetched \"%s\" from [%d %v]",
			update.Update.Message.Text,
			update.Message.From.ID,
			update.Message.From.Username),
		)
		currency, currencyErr := userService.GetInfoAboutCurrency(ctx, update.Update)
		var err error
		if currencyErr == botErrors.CurrencyNotFound {
			err = update.Answer(
				tg.HTML.Text("No such currency:(\ntry another")).ParseMode(tg.HTML).DoVoid(ctx)
		} else if currencyErr == botErrors.InvalidInput {
			err = update.Answer("Invalid command input. Check /help and try again\n").ParseMode(tg.HTML).DoVoid(ctx)
		} else {
			err = update.Answer(
				tg.HTML.Text("Currency info:",
					tg.HTML.Code(fmt.Sprintf("IsCrypto: %b\nValute name: %s\nUSDPrice: %f\n"+
						"Trade start date:%v\nTrade end date: %v",
						currency.IsCrypto, currency.Name, currency.PriceUSD, currency.TradeStart, currency.TradeEnd)))).
				ParseMode(tg.HTML).DoVoid(ctx)
		}
		botLogger.Info().Msg(fmt.Sprintf("[CurrencyInfo] Done \"%s\" from [%d %v].",
			update.Update.Message.Text,
			update.Message.From.ID,
			update.Update.Chat().Username))
		return err
	}, tgb.Command("currencyinfo", tgb.WithCommandIgnoreCase(true)))
}

func ClearStatsCommand(router *tgb.Router, service service.UserService) {
	botLogger := logger.GetLogger()
	router.Message(func(ctx context.Context, update *tgb.MessageUpdate) error {
		botLogger.Info().Msg(fmt.Sprintf("[Clear] Fetched \"%s\" from [%d %v]",
			update.Update.Message.Text,
			update.Message.From.ID,
			update.Update.Chat().Username),
		)
		err := service.ClearStats(ctx, update.Message.From.ID)
		if err != nil {
			botLogger.Warn().Err(err).Msg("[ClearStats] Fail")
		}
		err = update.Answer(tg.HTML.Text("Done!")).ParseMode(tg.HTML).DoVoid(ctx)
		botLogger.Info().Msg(fmt.Sprintf("[Clear] Done \"%s\" from [%d %v].",
			update.Update.Message.Text,
			update.Message.From.ID,
			update.Update.Chat().Username),
		)
		return err
	}, tgb.Command("clearstats", tgb.WithCommandIgnoreCase(true)))
}

func UnknownCommand(router *tgb.Router) {
	botLogger := logger.GetLogger()
	router.Message(func(ctx context.Context, update *tgb.MessageUpdate) error {
		botLogger.Warn().Msg(fmt.Sprintf("[Unknown] Fetched \"%s\" from [%d %v]",
			update.Update.Message.Text,
			update.Message.From.ID,
			update.Update.Chat().Username))
		err := update.Answer(
			tg.HTML.Text("Unknown command. Try another or /help")).
			ParseMode(tg.HTML).DoVoid(ctx)
		return err
	})
}
