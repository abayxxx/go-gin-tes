package config

import (
	"go-gin/app/controller"
	"go-gin/app/repository"
	"go-gin/app/service"
)

type InitializationUser struct {
	authRepo repository.AuthRepository
	authSvc  service.AuthService
	AuthCtrl controller.AuthController
}

type InitializationProduct struct {
	productRepo repository.ProductRepository
	productSvc  service.ProductService
	ProductCtrl controller.ProductController
}

type InitializationOrder struct {
	OrderRepo repository.OrderRepository
	OrderSvc  service.OrderService
	OrderCtrl controller.OrderController
}

func NewInitializationUser(authRepo repository.AuthRepository,
	authService service.AuthService,
	authCtrl controller.AuthController) *InitializationUser {
	return &InitializationUser{
		authRepo: authRepo,
		authSvc:  authService,
		AuthCtrl: authCtrl,
	}
}

func NewInitializationProduct(productRepo repository.ProductRepository,
	productService service.ProductService,
	productCtrl controller.ProductController) *InitializationProduct {
	return &InitializationProduct{
		productRepo: productRepo,
		productSvc:  productService,
		ProductCtrl: productCtrl,
	}
}

func NewInitializationOrder(orderRepo repository.OrderRepository,
	orderService service.OrderService,
	orderCtrl controller.OrderController) *InitializationOrder {
	return &InitializationOrder{
		OrderRepo: orderRepo,
		OrderSvc:  orderService,
		OrderCtrl: orderCtrl,
	}
}
