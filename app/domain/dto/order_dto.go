package dto

type OrderHistoryResponse struct {
	ID          uint64  `json:"id"`
	ProductID   uint64  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Total       float64 `json:"total"`
	Status      string  `json:"status"`
}

type ShoppingCartRequest struct {
	ProductID uint64 `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type ShoppingCartResponse struct {
	ID                 uint64  `json:"id"`
	ProductID          uint64  `json:"product_id"`
	ProductName        string  `json:"product_name"`
	ProductPhoto       string  `json:"product_photo"`
	ProductDescription string  `json:"product_description"`
	Price              float64 `json:"price"`
	Quantity           int     `json:"quantity"`
	Total              float64 `json:"total"`
}

type OrderRequest struct {
	ProductID uint64 `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
