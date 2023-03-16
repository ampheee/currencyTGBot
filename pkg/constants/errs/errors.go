package botErrors

import (
	"errors"
)

const (
	InvalidInput     = 1
	CurrencyNotFound = 2
)

var (
	NoDBURL        = errors.New("dburl not found in .env")
	NoTGToken      = errors.New("telegram token not found in .env")
	NoCoinAPIKey   = errors.New("coinApiKey not found in .env")
	NoSuchCurrency = errors.New("no such currency")
)
