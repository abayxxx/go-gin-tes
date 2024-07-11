package service

import (
	"go-gin/app/domain/dto"
	"go-gin/app/domain/model"
)

type AuthService interface {
	Login(request *dto.LoginRequest) (response *dto.LoginResponse, err error)
	Register(request *dto.RegisterRequest) (response bool, err error)
	GetUserById(userId uint) (response model.User, err error)
	TopUpWallet(userId uint, request *dto.TopUpWalletRequest) (response bool, err error)
	SSOAuth(request *dto.SSOAuthRequest) (response *dto.LoginResponse, err error)
}
