package repository

import (
	"go-gin/app/domain/dto"
	"go-gin/app/domain/model"
)

type AuthRepository interface {
	Login(email string, password string) (response model.User, err error)
	Register(request *dto.RegisterRequest) (response bool, err error)
	GetUserById(userId uint) (response model.User, err error)
	TopUpWallet(userId uint, request *dto.TopUpWalletRequest) (response bool, err error)
	SSOAuth(request *dto.SSOAuthRequest) (response model.User, err error)
}
