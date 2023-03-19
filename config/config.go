package config

import (
	"_entryTask/pkg/logger"
	"github.com/spf13/viper"
)

type Config struct {
	Postgres struct {
		User     *string `json:"user"`
		Password *string `json:"password"`
		Host     *string `json:"host"`
		Port     *string `json:"port"`
		DbName   *string `json:"dbname"`
	} `json:"postgres"`
	Tokens struct {
		TgToken   *string `json:"tgtoken"`
		CoinToken *string `json:"cointoken"`
	} `json:"tokens"`
}

func LoadConfig() *viper.Viper {
	log := logger.GetLogger()
	v := viper.New()
	v.AddConfigPath("../config")
	v.SetConfigName("config")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to read config")
	}
	log.Info().Msg("Config loaded successfully")
	return v
}

func ParseConfig(v *viper.Viper) *Config {
	log := logger.GetLogger()
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to decode config into struct")
	}
	log.Info().Msg("Config parsed successfully")
	return &c
}
