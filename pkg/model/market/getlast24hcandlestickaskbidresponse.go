package market

import "github.com/shopspring/decimal"

type GetLast24hCandlestickAskBidResponse struct {
	Status string             `json:"status"`
	Ch     string             `json:"ch"`
	Ts     int64              `json:"ts"`
	Tick   *CandlestickAskBid `json:"tick"`
}
type CandlestickAskBid struct {
	Id            int64             `json:"id"`
	Amount        decimal.Decimal   `json:"amount"`
	Open          decimal.Decimal   `json:"open"`
	Close         decimal.Decimal   `json:"close"`
	Count         int64             `json:"count"`
	High          decimal.Decimal   `json:"high"`
	Low           decimal.Decimal   `json:"low"`
	Vol           decimal.Decimal   `json:"vol"`
	TradeTurnover decimal.Decimal   `json:"trade_turnover"`
	Ts            int64             `json:"ts"`
	Bid           []decimal.Decimal `json:"bid"`
	Ask           []decimal.Decimal `json:"ask"`
}
