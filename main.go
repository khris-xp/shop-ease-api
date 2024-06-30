package main

import (
	"github.com/khris-xp/shop-ease-api/configs"
	"github.com/khris-xp/shop-ease-api/database"
	"github.com/khris-xp/shop-ease-api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	database.ConnectDB()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	routes.AuthRoutes(e)
	port := configs.EnvPort()
	e.Logger.Fatal(e.Start(":" + port))
}
