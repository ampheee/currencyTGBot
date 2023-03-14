package botErrors

import (
	"errors"
)

var (
	NoDBURL        = errors.New("dburl not found in .env")
	NoTGToken      = errors.New("telegram token not found in .env")
	NoCoinAPIKey   = errors.New("coinApiKey not found in .env")
	NoSuchCurrency = errors.New("no such currency")
)

var (
	NoDBURLMSG        = "i couldn't find the database url. Check, is dburl exist"
	NoTGTokenMSG      = "i couldn't find the telegram token. Check, is token in .env"
	NoCoinAPIKeyMSG   = "i couldn't find the coinApi key. Check, is key in .env"
	NoSuchCurrencyMSG = "oh, there`s no such currency. Try another :("
)
