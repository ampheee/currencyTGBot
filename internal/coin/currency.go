package coin

import (
	"_entryTask/config"
	"_entryTask/pkg/logger"
	"encoding/json"
	"gopkg.in/resty.v0"
)

func GetCurrencyInfo(config config.Config, coin string) *Currency {
	var currency []Currency
	botLogger := logger.GetLogger()
	resp, err := resty.R().
		SetHeader("X-CoinAPI-key", config.CoinAPIToken).
		Get("https://rest.coinapi.io/v1/assets/" + coin)
	if err != nil {
		botLogger.Fatal().Msg("Cant get request from api")
	}
	err = json.Unmarshal(resp.Body, &currency)
	if err != nil {
		botLogger.Warn().Err(err).Msg("Unable unmarshal to currency")
	}
	if string(resp.Body) == "[]" {
		return nil
	}
	return &currency[0]
}
