package model

type GetNumber struct {
	Status bool `json:"status"`
	Data   struct {
		OrderID string `json:"order_id"`
		Number  string `json:"number"`
		Status  string `json:"status"`
	} `json:"data"`
}
