package order

type PlaceOrdersResponse struct {
	Status  string        `json:"status"`
	Ts      int64         `json:"ts"`
	Success []OrderResult `json:"success"`
	Errors  []OrderError  `json:"errors"`
}

type OrderError struct {
	Index   int    `json:"index"`
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}
