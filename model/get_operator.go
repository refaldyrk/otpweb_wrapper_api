package model

type GetOperator struct {
	Status bool             `json:"status"`
	Data   map[int][]string `json:"data"`
}
