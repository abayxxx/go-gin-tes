//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"go-gin/app/controller"
	"go-gin/app/repository"
	"go-gin/app/service"
	"go-gin/app/service/jwt"
)

var DbConnection = wire.NewSet(
	ConnectToDB,
)

var AuthSet = wire.NewSet(
	repository.NewAuthRepositoryImpl,
	wire.Bind(new(repository.AuthRepository), new(*repository.AuthRepositoryImpl)),
	jwt.NewJwtService,
	service.NewAuthServiceImpl,
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),
	controller.NewAuthController,
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
)

var ProductSet = wire.NewSet(
	repository.NewProductRepositoryImpl,
	wire.Bind(new(repository.ProductRepository), new(*repository.ProductRepositoryImpl)),
	jwt.NewJwtService,
	service.NewProductServiceImpl,
	wire.Bind(new(service.ProductService), new(*service.ProductServiceImpl)),
	controller.NewProductController,
	wire.Bind(new(controller.ProductController), new(*controller.ProductControllerImpl)),
)

var OrderSet = wire.NewSet(
	repository.NewOrderRepositoryImpl,
	wire.Bind(new(repository.OrderRepository), new(*repository.OrderRepositoryImpl)),
	repository.NewAuthRepositoryImpl,
	wire.Bind(new(repository.AuthRepository), new(*repository.AuthRepositoryImpl)),
	jwt.NewJwtService,
	service.NewOrderServiceImpl,
	wire.Bind(new(service.OrderService), new(*service.OrderServiceImpl)),
	controller.NewOrderController,
	wire.Bind(new(controller.OrderController), new(*controller.OrderControllerImpl)),
)

func InitializeAuthController() *InitializationUser {
	wire.Build(
		NewInitializationUser,
		DbConnection,
		AuthSet,
	)
	return nil
}

func InitializeProductController() *InitializationProduct {
	wire.Build(
		NewInitializationProduct,
		DbConnection,
		ProductSet,
	)
	return nil
}

func InitializeOrderController() *InitializationOrder {
	wire.Build(
		NewInitializationOrder,
		DbConnection,
		OrderSet,
	)
	return nil
}
