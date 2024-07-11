// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"github.com/google/wire"
	"go-gin/app/controller"
	"go-gin/app/repository"
	"go-gin/app/service"
	"go-gin/app/service/jwt"
)

// Injectors from injector.go:

func InitializeAuthController() *InitializationUser {
	db := ConnectToDB()
	authRepositoryImpl := repository.NewAuthRepositoryImpl(db)
	jwtService := jwt.NewJwtService()
	authServiceImpl := service.NewAuthServiceImpl(authRepositoryImpl, jwtService)
	authControllerImpl := controller.NewAuthController(authServiceImpl)
	initializationUser := NewInitializationUser(authRepositoryImpl, authServiceImpl, authControllerImpl)
	return initializationUser
}

func InitializeProductController() *InitializationProduct {
	db := ConnectToDB()
	productRepositoryImpl := repository.NewProductRepositoryImpl(db)
	jwtService := jwt.NewJwtService()
	productServiceImpl := service.NewProductServiceImpl(productRepositoryImpl, jwtService)
	productControllerImpl := controller.NewProductController(productServiceImpl)
	initializationProduct := NewInitializationProduct(productRepositoryImpl, productServiceImpl, productControllerImpl)
	return initializationProduct
}

func InitializeOrderController() *InitializationOrder {
	db := ConnectToDB()
	orderRepositoryImpl := repository.NewOrderRepositoryImpl(db)
	authRepositoryImpl := repository.NewAuthRepositoryImpl(db)
	jwtService := jwt.NewJwtService()
	orderServiceImpl := service.NewOrderServiceImpl(orderRepositoryImpl, authRepositoryImpl, jwtService)
	orderControllerImpl := controller.NewOrderController(orderServiceImpl)
	initializationOrder := NewInitializationOrder(orderRepositoryImpl, orderServiceImpl, orderControllerImpl)
	return initializationOrder
}

// injector.go:

var DbConnection = wire.NewSet(
	ConnectToDB,
)

var AuthSet = wire.NewSet(repository.NewAuthRepositoryImpl, wire.Bind(new(repository.AuthRepository), new(*repository.AuthRepositoryImpl)), jwt.NewJwtService, service.NewAuthServiceImpl, wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)), controller.NewAuthController, wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)))

var ProductSet = wire.NewSet(repository.NewProductRepositoryImpl, wire.Bind(new(repository.ProductRepository), new(*repository.ProductRepositoryImpl)), jwt.NewJwtService, service.NewProductServiceImpl, wire.Bind(new(service.ProductService), new(*service.ProductServiceImpl)), controller.NewProductController, wire.Bind(new(controller.ProductController), new(*controller.ProductControllerImpl)))

var OrderSet = wire.NewSet(repository.NewOrderRepositoryImpl, wire.Bind(new(repository.OrderRepository), new(*repository.OrderRepositoryImpl)), repository.NewAuthRepositoryImpl, wire.Bind(new(repository.AuthRepository), new(*repository.AuthRepositoryImpl)), jwt.NewJwtService, service.NewOrderServiceImpl, wire.Bind(new(service.OrderService), new(*service.OrderServiceImpl)), controller.NewOrderController, wire.Bind(new(controller.OrderController), new(*controller.OrderControllerImpl)))
