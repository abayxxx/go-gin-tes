package model

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Photo       string  `json:"photo"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	BaseModel
}
