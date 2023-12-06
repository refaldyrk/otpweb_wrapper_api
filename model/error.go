package model

type Error struct {
	Status   bool   `json:"status"`
	Messages string `json:"messages"`
}
