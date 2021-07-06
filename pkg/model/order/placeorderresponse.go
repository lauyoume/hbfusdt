package order

type PlaceOrderResponse struct {
	Status  string      `json:"status"`
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Ts      int64       `json:"ts"`
	Data    OrderResult `json:"data"`
}

type OrderResult struct {
	Index         int    `json:"index"`
	OrderId       int64  `json:"order_id"`
	ClientOrderId int64  `json:"client_order_id"`
	OrderIdStr    string `json:"order_id_str"`
}
