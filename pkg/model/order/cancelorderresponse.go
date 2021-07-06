package order

type CancelOrderResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		Errors []struct {
			OrderId    string `json:"order_id"`
			ErrCode    string `json:"err_code"`
			ErrMessage string `json:"err_msg"`
		} `json:"errors"`
		Successes string `json:"successes"`
	} `json:"data"`
}
