package model

type GetService struct {
	Status bool `json:"status"`
	Data   []struct {
		ServiceID   string `json:"service_id"`
		ServiceName string `json:"service_name"`
		Cost        int64  `json:"cost"`
		Count       int64  `json:"count"`
	} `json:"data"`
}
