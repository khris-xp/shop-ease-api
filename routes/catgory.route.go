package routes

import (
	"github.com/khris-xp/shop-ease-api/controllers"
	"github.com/khris-xp/shop-ease-api/middlewares"
	"github.com/khris-xp/shop-ease-api/repositories"
	"github.com/labstack/echo/v4"
)

func CategoryRoutes(e *echo.Echo) {
	categoryRepo := repositories.NewCategoryRepository()
	categoryController := controllers.NewCategoryController(categoryRepo)

	category := e.Group("/api/categories")

	category.POST("", categoryController.CreateCategory, middlewares.AuthMiddleware)
	category.GET("", categoryController.GetCategories)
	category.GET("/:id", categoryController.GetCategoryByID)
	category.PUT("/:id", categoryController.UpdateCategory, middlewares.AuthMiddleware)
}
