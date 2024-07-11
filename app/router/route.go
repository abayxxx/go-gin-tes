package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-gin/app/middleware"
	"go-gin/app/service"
	"go-gin/config"
	"go-gin/docs"
)

func InitializeRouteV1(user config.InitializationUser, products config.InitializationProduct, order config.InitializationOrder) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = "/api/v1"
	authService := service.AuthServiceImpl{}

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		auth.POST("/login", user.AuthCtrl.Login)
		auth.POST("/register", user.AuthCtrl.Register)
		auth.POST("/sso/login", user.AuthCtrl.SSOLogin)
		auth.GET("/sso/callback", user.AuthCtrl.SSOCallback)

		product := api.Group("/product")
		product.GET("/", products.ProductCtrl.GetAll)
		product.GET("/:id", products.ProductCtrl.GetById)
		product.POST("/", products.ProductCtrl.StoreProduct)
		product.PUT("/:id", products.ProductCtrl.UpdateProduct)
		product.DELETE("/:id", products.ProductCtrl.DeleteProduct)

		orderRoute := api.Group("/order").Use(middleware.JwtMiddleware(&authService))
		orderRoute.GET("/:user_id", order.OrderCtrl.GetAllOrderUser)
		orderRoute.GET("/:user_id/:order_id", order.OrderCtrl.GetDetailOrderUser)
		orderRoute.POST("/cart/:user_id", order.OrderCtrl.StoreShoppingCart)
		orderRoute.GET("/cart/:user_id", order.OrderCtrl.GetShoppingCartList)
		orderRoute.POST("/checkout/:user_id", order.OrderCtrl.StoreOrder)

		userRoute := api.Group("/user")
		userRoute.POST("/top-up/:user_id", user.AuthCtrl.TopUpWallet)

		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	}

	return router
}
