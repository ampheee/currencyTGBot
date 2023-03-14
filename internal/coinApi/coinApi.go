package coinApi

import (
	botErrors "_entryTask/pkg/constants/errors"
	"_entryTask/pkg/logger"
	"gopkg.in/resty.v0"
)

func get() {
	botLogger := logger.GetLogger()
	resp, err := resty.R().
		SetHeader("X-CoinAPI-key", "44060329-D324-4E46-A8AE-72518CDF0A53").
		Get("https://rest.coinapi.io/v1/assets/bt121")
	if err != nil {
		botLogger.Fatal().Msg("Cant get request from api")
	}
	if string(resp.Body) == "[]" {
		botLogger.Err(botErrors.NoSuchCurrency).Msg("Coin does not exist")
	}
}
