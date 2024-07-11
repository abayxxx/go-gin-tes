package controller

import "github.com/gin-gonic/gin"

type ProductController interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	StoreProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}
