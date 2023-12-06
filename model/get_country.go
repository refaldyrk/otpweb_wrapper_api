package model

type GetCoutry struct {
	Status bool `json:"status"`
	Data   []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}
