package common

type GetMarketStatusResponse struct {
	Status string       `json:"status"`
	Data   MarketStatus `json:"data"`
	Ts     int64        `json:"ts"`
}

type MarketStatus struct {
	Heartbeat                       int   `json:"heartbeat"`
	SwapHeartbeat                   int   `json:"swap_heartbeat"`
	EstimatedRecoveryTime           int64 `json:"estimated_recovery_time"`
	SwapEstimatedRecoveryTime       int64 `json:"swap_estimated_recovery_time"`
	OptionHeartbeat                 int   `json:"option_heartbeat"`
	OptionEstimatedRecoveryTime     int   `json:"option_estimated_recovery_time"`
	LinearSwapHeartbeat             int   `json:"linear_swap_heartbeat"`
	LinearSwapEstimatedRecoveryTime int64 `json:"linear_swap_estimated_recovery_time"`
}
