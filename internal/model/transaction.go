package model

type Transaction struct {
	ID     int     `json:"id"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}
