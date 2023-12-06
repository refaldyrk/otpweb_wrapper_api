package model

type GetBalance struct {
	Status bool        `json:"status"`
	Data   dataBalance `json:"data"`
}

type dataBalance struct {
	Balance string `json:"balance"`
}
