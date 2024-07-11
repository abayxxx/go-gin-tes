package controller

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-gin/app/constant"
	"go-gin/app/domain/dto"
	pkg "go-gin/app/helper"
	"go-gin/app/service"
	ssoConfig "go-gin/app/sso_config"
	"net/http"
	"strconv"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

// create random string 16 char
var oauthStateString, _ = pkg.GenerateRandomString(16)

func NewAuthController(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

// @BasePath /api/v1

// Login PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Param request body dto.LoginRequest true "Login Request"
// @Accept json
// @Produce json
// @Success 200 {object} dto.LoginResponse
// @Router /auth/login [post]
func (a AuthControllerImpl) Login(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request dto.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	response, err := a.AuthService.Login(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

func (a AuthControllerImpl) Register(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request dto.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	response, err := a.AuthService.Register(&request)
	if err != nil {
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

// TopUpWallet is a function to handle top up wallet request
func (a AuthControllerImpl) TopUpWallet(c *gin.Context) {
	defer pkg.PanicHandler(c)

	//get id from path
	id := c.Param("user_id")

	//convert id to uint64
	idUint64, err := strconv.ParseUint(id, 10, 64)

	var request dto.TopUpWalletRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	response, err := a.AuthService.TopUpWallet(uint(idUint64), &request)
	if err != nil {
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

// SSOLogin is a function to handle sso login request
func (a AuthControllerImpl) SSOLogin(c *gin.Context) {
	defer pkg.PanicHandler(c)

	url := ssoConfig.GoogleOauthConfig.AuthCodeURL(oauthStateString)
	c.JSON(http.StatusOK, gin.H{"url": url})
}

// SSOCallback is a function to handle sso callback request
func (a AuthControllerImpl) SSOCallback(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		log.Println("invalid oauth state")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid oauth state"})
		return
	}

	code := c.Query("code")
	token, err := ssoConfig.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println("code exchange failed: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "code exchange failed"})
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Println("failed getting user info: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed getting user info"})
		return
	}
	defer response.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		log.Println("failed to decode user info: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to decode user info"})
		return
	}

	var request dto.SSOAuthRequest
	request.Email = userInfo["email"].(string)
	request.Name = userInfo["name"].(string)

	//call service sso auth
	resp, err := a.AuthService.SSOAuth(&request)
	if err != nil {
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, resp))
}
