package market

import "github.com/shopspring/decimal"

type GetDepthResponse struct {
	Status string `json:"status"`
	Ch     string `json:"ch"`
	Ts     int64  `json:"ts"`
	Tick   *Depth `json:"tick"`
}

type Depth struct {
	Ch        string              `json:"ch"`
	Id        int64               `json:"id"`
	Mrid      int64               `json:"mrid"`
	Timestamp int64               `json:"ts"`
	Version   int64               `json:"version"`
	Bids      [][]decimal.Decimal `json:"bids"`
	Asks      [][]decimal.Decimal `json:"asks"`
}
