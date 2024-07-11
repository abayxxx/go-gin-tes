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

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(authService service.OrderService) *OrderControllerImpl {
	return &OrderControllerImpl{
		OrderService: authService,
	}
}

func (o *OrderControllerImpl) GetAllOrderUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	//get id from path
	id := c.Param("user_id")

	//convert id to uint64
	idUint64, err := strconv.ParseUint(id, 10, 64)

	response, err := o.OrderService.GetAllOrderUser(idUint64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

func (o *OrderControllerImpl) GetDetailOrderUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	//get id from path
	id := c.Param("user_id")
	orderId := c.Param("order_id")

	//convert id to uint64
	idUint64, err := strconv.ParseUint(id, 10, 64)
	orderIdUint64, err := strconv.ParseUint(orderId, 10, 64)

	response, err := o.OrderService.GetDetailOrderUser(idUint64, orderIdUint64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

func (o *OrderControllerImpl) StoreShoppingCart(c *gin.Context) {

	defer pkg.PanicHandler(c)
	//get id from path
	id := c.Param("user_id")

	//convert id to uint64
	idUint64, err := strconv.ParseUint(id, 10, 64)

	var request dto.ShoppingCartRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	response, err := o.OrderService.StoreShoppingCart(idUint64, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

func (o *OrderControllerImpl) GetShoppingCartList(c *gin.Context) {
	defer pkg.PanicHandler(c)

	//get id from path
	id := c.Param("user_id")

	//convert id to uint64
	idUint64, err := strconv.ParseUint(id, 10, 64)

	response, err := o.OrderService.GetShoppingCartList(idUint64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

func (o *OrderControllerImpl) StoreOrder(c *gin.Context) {
	defer pkg.PanicHandler(c)

	//get id from path
	id := c.Param("user_id")

	//convert id to uint64
	idUint64, err := strconv.ParseUint(id, 10, 64)

	var request dto.OrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	response, err := o.OrderService.StoreOrder(idUint64, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}
