package common

import "github.com/shopspring/decimal"

type GetSymbolsResponse struct {
	Status string   `json:"status"`
	Data   []Symbol `json:"data"`
}

type Symbol struct {
	Symbol            string          `json:"symbol"`
	ContractCode      string          `json:"contract_code"`
	ContractType      string          `json:"contract_type"`
	ContractSize      decimal.Decimal `json:"contract_size"`
	PriceTick         decimal.Decimal `json:"price_tick"`
	DeliveryTime      string          `json:"delivery_time"`
	CreateDate        string          `json:"create_date"`
	ContractStatus    int             `json:"contract_status"`
	SupportMarginMode string          `json:"support_margin_mode"`
}
