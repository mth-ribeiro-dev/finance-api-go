package model

type Transaction struct {
	ID          int      `json:"id"`
	Type        string   `json:"type"`
	Amount      float64  `json:"amount"`
	Category    string   `json:"category"`
	Date        DateOnly `json:"date"`
	Description string   `json:"description"`
	UserID      int      `json:"user_id"`
}
