package common

type GetTimestampResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
}
