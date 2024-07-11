package controller

import "github.com/gin-gonic/gin"

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	SSOLogin(ctx *gin.Context)
	SSOCallback(ctx *gin.Context)
	TopUpWallet(ctx *gin.Context)
}
