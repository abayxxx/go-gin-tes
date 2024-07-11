package service

import "go-gin/app/domain/dto"

type OrderService interface {
	GetAllOrderUser(userId uint64) (response []dto.OrderHistoryResponse, err error)
	GetDetailOrderUser(userId, orderId uint64) (response dto.OrderHistoryResponse, err error)
	StoreShoppingCart(userId uint64, request dto.ShoppingCartRequest) (response bool, err error)
	GetShoppingCartList(userId uint64) (response []dto.ShoppingCartResponse, err error)
	StoreOrder(userId uint64, request dto.OrderRequest) (response bool, err error)
}
