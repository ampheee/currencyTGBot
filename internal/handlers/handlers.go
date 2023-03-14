package handlers

import (
	"_entryTask/pkg/logger"
	"context"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
)

func InitHandlers(router *tgb.Router) {
	StartCommand(router)
	HelpCommand(router)
	StatsCommand(router)
	ExchangeCommand(router)
	CurrencyInfoCommand(router)
	ClearStatsCommand(router)
	UnknownCommand(router)
}

func StartCommand(router *tgb.Router) {
	botLogger := logger.GetLogger()
	router.Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
		botLogger.Info().Msg(fmt.Sprintf("[Start] Fetched \"%s\" from [%d %v]", msg.Update.Message.Text,
			msg.Update.ID, msg.Update.Chat().Username))
		err := msg.Answer(
			tg.HTML.Text(
				tg.HTML.Line("It`s entry-task product"),
				tg.HTML.Bold("\n 💸 I`m currency bot 💸"),
				"",
				tg.HTML.Italic("🚀 Powered by ampheee 🚀",
					"\n",
					tg.HTML.Underline(tg.HTML.Link("my tg", "github.com/ampheee")))),
		).ParseMode(tg.HTML).DoVoid(ctx)
		return err
	}, tgb.Command("start", tgb.WithCommandIgnoreCase(true)))
}

func HelpCommand(router *tgb.Router) {
	botLogger := logger.GetLogger()
	router.Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
		botLogger.Info().Msg(fmt.Sprintf("[Start] Fetched \"%s\" from [%d %v]", msg.Update.Message.Text,
			msg.Update.ID, msg.Update.Chat().Username))
		err := msg.Answer(
			tg.HTML.Text(
				tg.HTML.Bold("All existing commands:\n"),
				tg.HTML.Code(
					fmt.Sprintf("%s: %s\n%s: %s\n%s:%s\n%s: %s\n%s: %s",
						"\t/start", tg.HTML.Bold("calling greetings with credits"),
						"\t/stats", tg.HTML.Bold("shows your statistic for all time"),
						"\t/exchange", tg.HTML.Bold("exchanges one currency into another"),
						"\t/currencyInfo", tg.HTML.Bold("get entered currency info from coinAPI"),
						"\t/clearstats", tg.HTML.Bold("clear all your stats\n"),
					)),
				tg.HTML.Text("-------------------------------------------------"+
					"--------------------------------------------------------------"),
				tg.HTML.Bold("Example usage:\n"),
				tg.HTML.Text("\t/exchange BTC USD")),
		).ParseMode(tg.HTML).DoVoid(ctx)
		return err
	}, tgb.Command("help", tgb.WithCommandIgnoreCase(true)))
}

func StatsCommand(router *tgb.Router) {
	botLogger := logger.GetLogger()
	router.Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
		botLogger.Info().Msg(fmt.Sprintf("[Stats] Fetched \"%s\" from [%d %v]", msg.Update.Message.Text,
			msg.Update.ID, msg.Update.Chat().Username))
		//TODO ParseUserStats
		err := msg.Answer(
			tg.HTML.Text(fmt.Sprintf("Your stats:\n"+
				"Total requests: %d", 1))).
			ParseMode(tg.HTML).DoVoid(ctx)
		return err
	}, tgb.Command("stats"))
}

func ExchangeCommand(router *tgb.Router) {
	botLogger := logger.GetLogger()
	router.Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
		botLogger.Info().Msg(fmt.Sprintf("[CurrencyInfo] Fetched \"%s\" from [%d %v]",
			msg.Update.Message.Text,
			msg.Update.ID,
			msg.Update.Chat().Username),
		)
		//TODO GetCurrencyInfo
		err := msg.Answer(
			tg.HTML.Text()).ParseMode(tg.HTML).DoVoid(ctx)
		botLogger.Info().Msg(fmt.Sprintf("[CurrencyInfo] Done \"%s\" from [%d %v].",
			msg.Update.Message.Text,
			msg.Update.ID,
			msg.Update.Chat().Username),
		)
		return err
	}, tgb.Command("exchange", tgb.WithCommandIgnoreCase(true)))
}

func CurrencyInfoCommand(router *tgb.Router) {
	botLogger := logger.GetLogger()
	router.Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
		botLogger.Info().Msg(fmt.Sprintf("[Exchange] Fetched \"%s\" from [%d %v]",
			msg.Update.Message.Text,
			msg.Update.ID,
			msg.Update.Chat().Username),
		)
		//TODO ExchangeCurrencies
		err := msg.Answer(
			tg.HTML.Text()).ParseMode(tg.HTML).DoVoid(ctx)
		botLogger.Info().Msg(fmt.Sprintf("[Exchange] Done \"%s\" from [%d %v].",
			msg.Update.Message.Text,
			msg.Update.ID,
			msg.Update.Chat().Username),
		)
		return err
	}, tgb.Command("currencyinfo", tgb.WithCommandIgnoreCase(true)))
}

func ClearStatsCommand(router *tgb.Router) {
	botLogger := logger.GetLogger()
	router.Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
		botLogger.Info().Msg(fmt.Sprintf("[Clear] Fetched \"%s\" from [%d %v]",
			msg.Update.Message.Text,
			msg.Update.ID,
			msg.Update.Chat().Username),
		)
		//TODO ClearHistory
		err := msg.Answer(tg.HTML.Text("Done!")).ParseMode(tg.HTML).DoVoid(ctx)
		botLogger.Info().Msg(fmt.Sprintf("[Clear] Done \"%s\" from [%d %v].",
			msg.Update.Message.Text,
			msg.Update.ID,
			msg.Update.Chat().Username),
		)
		return err
	}, tgb.Command("clearstats", tgb.WithCommandIgnoreCase(true)))
}

func UnknownCommand(router *tgb.Router) {
	botLogger := logger.GetLogger()
	router.Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
		botLogger.Warn().Msg(fmt.Sprintf("[Unknown] Fetched \"%s\" from [%d %v]", msg.Update.Message.Text,
			msg.Update.ID, msg.Update.Chat().Username))
		err := msg.Answer(
			tg.HTML.Text("Unknown command. Try another or /help")).
			ParseMode(tg.HTML).DoVoid(ctx)
		return err
	})
}
