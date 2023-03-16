package coin

import "time"

type Currency struct {
	Name       string    `json:"name"`
	IsCrypto   int       `json:"type_is_crypto"`
	TradeStart time.Time `json:"data_trade_start"`
	TradeEnd   time.Time `json:"data_trade_end"`
	PriceUSD   float64   `json:"price_usd"`
}
