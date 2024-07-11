package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"go-gin/app/domain/dto"
)

type JwtService interface {
	GenerateToken(memberID uint) (dto.LoginResponse, error)
	ValidateToken(token string) (*jwt.Token, error)
	IsTokenExpired(token string) bool
}
