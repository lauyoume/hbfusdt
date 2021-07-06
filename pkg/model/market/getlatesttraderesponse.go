package market

import "github.com/shopspring/decimal"

type GetLatestTradeResponse struct {
	Status string     `json:"status"`
	Ch     string     `json:"ch"`
	Ts     int64      `json:"ts"`
	Tick   *TradeTick `json:"tick"`
}
type TradeTick struct {
	Id   int64 `json:"id"`
	Ts   int64 `json:"ts"`
	Data []struct {
		Amount        decimal.Decimal `json:"amount"`
		Direction     string          `json:"direction"`
		Id            decimal.Decimal `json:"id"`
		Price         decimal.Decimal `json:"price"`
		Ts            int64           `json:"ts"`
		Quantity      decimal.Decimal `json:"quantity"`
		ContractCode  string          `json:"contract_code"`
		TradeTurnover decimal.Decimal `json:"trade_turnover"`
	}
}
