package repository

import (
	"errors"
	"fmt"
	"go-gin/app/domain/dto"
	"go-gin/app/domain/model"
	pkg "go-gin/app/helper"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

// NewAuthRepositoryImpl is a constructor function to create NewAuthRepositoryImpl
func NewAuthRepositoryImpl(db *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{db}
}

func (m *AuthRepositoryImpl) Login(email string, password string) (response model.User, err error) {
	users := model.User{}
	err = m.db.Table("users").Where("email = ? AND deleted_at IS NULL", email).First(&users).Error

	if err != nil {
		return users, errors.New(fmt.Sprintf("Email or Password is wrong"))
	}
	return users, nil

}

func (m *AuthRepositoryImpl) Register(request *dto.RegisterRequest) (response bool, err error) {
	err = m.db.Table("users").Create(&request).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *AuthRepositoryImpl) GetUserById(userId uint) (response model.User, err error) {
	users := model.User{}
	err = m.db.Table("users").Where("id = ? AND deleted_at IS NULL", userId).First(&users).Error
	if err != nil {
		return users, errors.New(fmt.Sprintf("User with id %d not found", userId))
	}
	return users, nil
}

func (m *AuthRepositoryImpl) TopUpWallet(userId uint, request *dto.TopUpWalletRequest) (response bool, err error) {
	user := model.UserWallet{}
	err = m.db.Table("user_wallets").Where("id = ? AND deleted_at IS NULL", userId).First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {

		//create new wallet
		user = model.UserWallet{
			UserID:  userId,
			Balance: float64(request.Amount),
		}
		err = m.db.Table("user_wallets").Create(&user).Error
		if err != nil {
			return false, err
		}

		return true, nil
	} else {
		user.Balance += float64(request.Amount)
		err = m.db.Table("user_wallets").Save(&user).Error
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (m *AuthRepositoryImpl) SSOAuth(request *dto.SSOAuthRequest) (response model.User, err error) {
	users := model.User{}
	err = m.db.Table("users").Where("email = ? AND deleted_at IS NULL", request.Email).First(&users).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		randPassword, _ := pkg.GenerateRandomString(10)
		//create new user
		user := model.User{
			Email:    request.Email,
			Password: randPassword,
			Name:     request.Name,
		}
		err = m.db.Table("users").Create(&user).Error
		if err != nil {
			return users, err
		}
		return user, nil
	}
	return users, nil
}
