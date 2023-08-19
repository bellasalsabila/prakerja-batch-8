package main

import (
	"coffeshop/config"
	"coffeshop/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.InitRoutes(e)
	config.InitDatabase()
	defer config.CloseDatabase()
	e.Start(":5530")
}
