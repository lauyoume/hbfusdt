package order

// 多个orderid使用,分割
type CancelOrderRequest struct {
	OrderId       string `json:"order_id,omitempty"`
	ClientOrderId string `json:"client_order_id,omitempty"`
	Symbol        string `json:"contract_code"`
}
