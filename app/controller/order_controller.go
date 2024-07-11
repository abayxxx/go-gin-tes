package controller

import "github.com/gin-gonic/gin"

type OrderController interface {
	GetAllOrderUser(ctx *gin.Context)
	GetDetailOrderUser(ctx *gin.Context)
	StoreShoppingCart(ctx *gin.Context)
	GetShoppingCartList(ctx *gin.Context)
	StoreOrder(ctx *gin.Context)
}
