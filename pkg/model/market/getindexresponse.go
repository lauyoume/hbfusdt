package market

import "github.com/shopspring/decimal"

type GetIndexResponse struct {
	Status string  `json:"status"`
	Ts     int64   `json:"ts"`
	Data   []Index `json:"data"`
}

type Index struct {
	ContractCode string          `json:"contract_code"`
	IndexPrice   decimal.Decimal `json:"index_price"`
	IndexTs      int64           `json:"index_ts"`
}
