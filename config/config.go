package config

import (
	errs "_entryTask/pkg/constants/errs"
	"_entryTask/pkg/logger"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBURL        string
	TGBotToken   string
	CoinAPIToken string
}

func GetConfig() Config {
	log := logger.GetLogger()
	err := godotenv.Load("../config/.env")
	if err != nil {
		log.Fatal().Err(err).Msg("Cant load your .env file/ Is he exist ?")
	}
	log.Info().Msg(".env File loaded successfully")
	dburl := os.Getenv("DBURL")
	if dburl == "" {
		log.Fatal().Err(errs.NoDBURL)
	}
	tgbottoken := os.Getenv("TGBOTTOKEN")
	if tgbottoken == "" {
		log.Fatal().Err(errs.NoTGToken)
	}
	coinapitoken := os.Getenv("COINAPIKEY")
	if coinapitoken == "" {
		log.Fatal().Err(errs.NoCoinAPIKey)
	}
	log.Info().Msg(".env Parsed")
	log.Info().Msg("Config loaded successfully")
	return Config{DBURL: dburl, CoinAPIToken: coinapitoken, TGBotToken: tgbottoken}
}
