package order

type PlaceOrderRequest struct {
	Symbol    string  `json:"contract_code"`
	Direction string  `json:"direction"`
	Offset    string  `json:"offset"`
	Price     float64 `json:"price,omitempty"`
	// Price            decimal.Decimal `json:"price,omitempty"`
	Volume         int     `json:"volume"`
	LeverRate      int     `json:"lever_rate"`
	OrderPriceType string  `json:"order_price_type"`
	ClientOrderId  string  `json:"client-order-id,omitempty"`
	TpTriggerPrice float64 `json:"tp_trigger_price,omitempty"`
	TpOrderPrice   float64 `json:"tp_order_price,omitempty"`
	// TpTriggerPrice   decimal.Decimal `json:"tp_trigger_price,omitempty"`
	// TpOrderPrice     decimal.Decimal `json:"tp_order_price,omitempty"`
	TpOrderPriceType string  `json:"tp_order_price_type,omitempty"`
	SlTriggerPrice   float64 `json:"sl_trigger_price,omitempty"`
	SlOrderPrice     float64 `json:"sl_order_price,omitempty"`
	// SlTriggerPrice   decimal.Decimal `json:"sl_trigger_price,omitempty"`
	// SlOrderPrice     decimal.Decimal `json:"sl_order_price,omitempty"`
	SlOrderPriceType string `json:"sl_order_price_type,omitempty"`
}

type PlaceOrderBatchRequest struct {
	OrdersData []PlaceOrderRequest `json:"orders_data"`
}

const (
	BUY  = "buy"
	SELL = "sell"

	OPEN  = "open"
	CLOSE = "close"
)

const (
	TYPE_LIMIT      = "limit"
	TYPE_POSTONLY   = "post_only"
	TYPE_OPPONENT   = "opponent"
	TYPE_OPTIMAL_5  = "optimal_5"
	TYPE_OPTIMAL_10 = "optimal_10"
	TYPE_OPTIMAL_20 = "optimal_20"

	// 闪电平仓 最优30挡
	TYPE_LIGHTNING     = "lightning"
	TYPE_LIGHTNING_IOC = "lightning_ioc"
	TYPE_LIGHTNING_FOK = "lightning_fok"
)

/*
开平方向
开多：买入开多(direction用buy、offset用open)
平多：卖出平多(direction用sell、offset用close)
开空：卖出开空(direction用sell、offset用open)
平空：买入平空(direction用buy、offset用close)
*/

// "limit":限价，"post_only":只做maker单，
// ioc:IOC订单，fok：FOK订单
// 这四种报价类型需要传价格，其他都不需要。
