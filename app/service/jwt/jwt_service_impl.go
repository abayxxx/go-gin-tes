package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-gin/app/domain/dto"
	"os"
	"strconv"
	"time"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

type jwtServiceImpl struct {
}

func NewJwtService() JwtService {
	return &jwtServiceImpl{}
}

func (j *jwtServiceImpl) GenerateToken(userId uint) (dto.LoginResponse, error) {
	var jwtExpTime = "60"
	//jwtExpTime = "15"
	//convert jwtExpTime to int64
	jwtExpTimeInt, _ := strconv.ParseInt(jwtExpTime, 10, 64)

	expiredTime := time.Now().Add(time.Minute * time.Duration(jwtExpTimeInt)) // expired in 30 min
	expiredTimeRt := time.Now().Add(time.Hour * time.Duration(1))             // expired in 1 hour
	claim := jwt.MapClaims{}

	claim["user_id"] = userId
	claim["refresh_token"] = false
	claim["exp"] = jwt.NewNumericDate(expiredTime).Unix() // expired at
	claim["iat"] = time.Now().Unix()                      // issued at

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return dto.LoginResponse{}, errors.New("error while signing token")
	}

	rtClaim := jwt.MapClaims{}

	rtClaim["user_id"] = userId
	rtClaim["refresh_token"] = true
	rtClaim["exp"] = jwt.NewNumericDate(expiredTimeRt).Unix()
	rtClaim["iat"] = time.Now().Unix()

	return dto.LoginResponse{
		TokenType:   "Bearer",
		ExpiresIn:   int(expiredTime.Unix()),
		AccessToken: signedToken,
	}, nil
}

func (j *jwtServiceImpl) ValidateToken(token string) (*jwt.Token, error) {
	newToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	return newToken, nil
}

func (j *jwtServiceImpl) IsTokenExpired(token string) bool {
	newToken, err := j.ValidateToken(token)
	if err != nil {
		return true
	}

	claims, ok := newToken.Claims.(jwt.MapClaims)
	if !ok {
		return true
	}

	expiredTime := time.Unix(int64(claims["exp"].(float64)), 0)
	return expiredTime.Before(time.Now())
}
