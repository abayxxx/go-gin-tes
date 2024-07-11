package repository

import (
	"go-gin/app/domain/dto"
	"go-gin/app/domain/model"
)

type OrderRepository interface {
	GetAllOrderUser(userId uint64) (response []model.UserOrderJournal, err error)
	GetDetailOrderUser(userId, orderId uint64) (response model.UserOrderJournal, err error)
	StoreShoppingCart(userId uint64, request dto.ShoppingCartRequest) (response bool, err error)
	GetShoppingCartList(userId uint64) (response []model.ShoppingCart, err error)
	StoreOrder(userId uint64, request dto.OrderRequest) (response bool, err error)
}
