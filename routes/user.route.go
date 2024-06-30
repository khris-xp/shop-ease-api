package routes

import (
	"github.com/khris-xp/shop-ease-api/controllers"
	"github.com/khris-xp/shop-ease-api/middlewares"
	"github.com/khris-xp/shop-ease-api/repositories"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	userRepo := repositories.NewUserRepository()
	authController := controllers.NewAuthController(userRepo)

	auth := e.Group("/api/auth")

	auth.POST("/register", authController.RegisterUser)
	auth.POST("/login", authController.LoginUser)
	auth.GET("/user", authController.GetUser, middlewares.AuthMiddleware)
	auth.GET("/refresh-token", authController.RefreshToken)
}
