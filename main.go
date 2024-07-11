package main

import (
	"github.com/joho/godotenv"
	"go-gin/app/router"
	"go-gin/config"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	initAuth := config.InitializeAuthController()
	initProduct := config.InitializeProductController()
	initOrder := config.InitializeOrderController()

	app := router.InitializeRouteV1(*initAuth, *initProduct, *initOrder)
	err := app.Run(":" + port)
	if err != nil {
		return
	}
}
