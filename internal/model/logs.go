package model

type Logs struct {
	ID     int64  `json:"id"`
	Char   string `json:"character" binding:"required"`
	Dialog string `json:"dialog" binding:"required"`
}
