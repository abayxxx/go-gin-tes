package model

type UserOrderJournal struct {
	ID              uint    `json:"id"`
	UserID          uint    `json:"user_id"`
	ProductID       uint    `json:"product_id"`
	Quantity        int     `json:"quantity"`
	Price           float64 `json:"price"`
	LastBalanceUser float64 `json:"last_balance_user"`
	Status          int     `json:"status"`
	BaseModel
	User    User    `gorm:"foreignKey:UserID"`
	Product Product `gorm:"foreignKey:ProductID"`
}
