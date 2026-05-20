package domain

type Wallet struct {
	UserID  int `json:"user_id"`
	Balance int    `json:"balance"`
}