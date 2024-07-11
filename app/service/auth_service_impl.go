package service

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-gin/app/domain/dto"
	"go-gin/app/domain/model"
	"go-gin/app/repository"
	"go-gin/app/service/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repository repository.AuthRepository
	JwtService jwt.JwtService
}

// NewAuthServiceImpl is a constructor function to create authServiceImpl
func NewAuthServiceImpl(repository repository.AuthRepository, jwtService jwt.JwtService) *AuthServiceImpl {
	return &AuthServiceImpl{repository, jwtService}
}

var validate = validator.New()

func (s *AuthServiceImpl) Login(request *dto.LoginRequest) (response *dto.LoginResponse, err error) {
	if err = validate.Struct(request); err != nil {
		return response, err
	}

	userNew, err := s.repository.Login(request.Email, request.Password)
	if err != nil {
		return response, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userNew.Password), []byte(request.Password))
	if err != nil {
		return response, errors.New(fmt.Sprintf("Email or Password is wrong"))
	}

	resultToken, err := s.JwtService.GenerateToken(userNew.ID)
	if err != nil {
		return response, err
	}

	//response
	response = &dto.LoginResponse{
		TokenType:   "Bearer",
		AccessToken: resultToken.AccessToken,
		ExpiresIn:   resultToken.ExpiresIn,
	}

	return response, nil
}

func (s *AuthServiceImpl) Register(request *dto.RegisterRequest) (response bool, err error) {
	if err = validate.Struct(request); err != nil {
		return response, err
	}

	//if s.repository.CheckEmailUnique(request.Email) {
	//	return response, errors.New(fmt.Sprint("Email already exists, Please register with different email"))
	//}

	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	request.Password = string(bytes)

	response, err = s.repository.Register(request)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (s *AuthServiceImpl) GetUserById(userId uint) (response model.User, err error) {
	user, err := s.repository.GetUserById(userId)
	if err != nil {
		return response, err
	}

	return user, nil
}

func (s *AuthServiceImpl) TopUpWallet(userId uint, request *dto.TopUpWalletRequest) (response bool, err error) {
	if err = validate.Struct(request); err != nil {
		return response, err
	}

	response, err = s.repository.TopUpWallet(userId, request)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (s *AuthServiceImpl) SSOAuth(request *dto.SSOAuthRequest) (response *dto.LoginResponse, err error) {
	if err = validate.Struct(request); err != nil {
		return response, err
	}

	user, err := s.repository.SSOAuth(request)
	if err != nil {
		return response, err
	}

	resultToken, err := s.JwtService.GenerateToken(user.ID)
	if err != nil {
		return response, err
	}

	//response
	response = &dto.LoginResponse{
		TokenType:   "Bearer",
		AccessToken: resultToken.AccessToken,
		ExpiresIn:   resultToken.ExpiresIn,
	}

	return response, nil
}
