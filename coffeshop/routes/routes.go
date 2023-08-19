package routes

import (
	"coffeshop/controller"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	api := e.Group("/api")

	// Public routes
	api.POST("/login", controller.Login)

	// Protected routes
	protect := api.Group("/protected")
	protect.Use(controller.JWTMiddleware()) // Middleware for JWT auth
	protect.GET("/data/:id", controller.GetData)
	protect.POST("/data", controller.CreateData)
	protect.PUT("/data/:id", controller.UpdateData)
	protect.DELETE("/data/:id", controller.DeleteData)
}
