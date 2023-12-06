package model

type GetStatus struct {
	Status bool `json:"status"`
	Data   struct {
		OrderId string `json:"order_id"`
		Number  string `json:"number"`
		Status  string `json:"status"`
		SMS     string `json:"SMS"`
	} `json:"data"`
}
