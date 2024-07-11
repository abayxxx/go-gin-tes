package model

type UserWallet struct {
	ID      uint    `json:"id"`
	UserID  uint    `json:"user_id"`
	Balance float64 `json:"balance"`
	BaseModel
}
