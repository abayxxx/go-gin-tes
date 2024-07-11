package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-gin/app/constant"
	"go-gin/app/domain/dto"
	pkg "go-gin/app/helper"
	"go-gin/app/service"
	"net/http"
	"strconv"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(authService service.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		ProductService: authService,
	}
}

func (p ProductControllerImpl) GetAll(c *gin.Context) {
	defer pkg.PanicHandler(c)

	response, err := p.ProductService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

func (p ProductControllerImpl) GetById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	//get id from path
	id := c.Param("id")

	//convert id to uint64
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	response, err := p.ProductService.GetById(idUint64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

func (p ProductControllerImpl) StoreProduct(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request dto.ProductRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	response, err := p.ProductService.StoreProduct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

func (p ProductControllerImpl) UpdateProduct(c *gin.Context) {
	defer pkg.PanicHandler(c)

	//get id from path
	id := c.Param("id")

	//convert id to uint64
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	var request dto.UpdateProductRequest
	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	response, err := p.ProductService.UpdateProduct(idUint64, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

func (p ProductControllerImpl) DeleteProduct(c *gin.Context) {
	defer pkg.PanicHandler(c)

	//get id from path
	id := c.Param("id")

	//convert id to uint64
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	response, err := p.ProductService.DeleteProduct(idUint64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}
