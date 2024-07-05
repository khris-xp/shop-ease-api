package routes

import (
	"github.com/khris-xp/shop-ease-api/controllers"
	"github.com/khris-xp/shop-ease-api/middlewares"
	"github.com/khris-xp/shop-ease-api/repositories"
	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo) {
	productRepo := repositories.NewProductRepository()
	productController := controllers.NewProductController(productRepo)

	product := e.Group("/api/products")

	product.POST("", productController.CreateProduct, middlewares.AuthMiddleware)
	product.GET("", productController.GetProducts)
	product.GET("/:id", productController.GetProductByID)
	product.PUT("/:id", productController.UpdateProduct, middlewares.AuthMiddleware)
}
