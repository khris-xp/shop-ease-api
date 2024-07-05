package controllers

import (
	"github.com/khris-xp/shop-ease-api/models"
	"github.com/khris-xp/shop-ease-api/repositories"
	"github.com/khris-xp/shop-ease-api/responses"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	ProductRepo *repositories.ProductRepository
}

func NewProductController(productRepo *repositories.ProductRepository) *ProductController {
	return &ProductController{ProductRepo: productRepo}
}

func (pc *ProductController) CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return responses.ProductErrorResponse(c, 400, "invalid request body")
	}

	message, err := pc.ProductRepo.CreateProduct(c.Request().Context(), product)
	if err != nil {
		return responses.ProductErrorResponse(c, 400, err.Error())
	}

	return responses.ProductSuccessResponse(c, 201, message, nil)
}

func (pc *ProductController) GetProducts(c echo.Context) error {
	products, err := pc.ProductRepo.GetProducts(c.Request().Context())
	if err != nil {
		return responses.ProductErrorResponse(c, 400, err.Error())
	}

	return responses.ProductSuccessResponse(c, 200, "success", products)
}

const invalidProductID = "invalid product ID"

func (pc *ProductController) GetProductByID(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return responses.ProductErrorResponse(c, 400, invalidProductID)
	}

	product, err := pc.ProductRepo.GetProductByID(c.Request().Context(), id)
	if err != nil {
		return responses.ProductErrorResponse(c, 400, err.Error())
	}

	return responses.ProductSuccessResponse(c, 200, "success", product)
}

func (pc *ProductController) UpdateProduct(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return responses.ProductErrorResponse(c, 400, invalidProductID)
	}

	var product models.Product
	if err := c.Bind(&product); err != nil {
		return responses.ProductErrorResponse(c, 400, "invalid request body")
	}

	message, err := pc.ProductRepo.UpdateProduct(c.Request().Context(), id, product)
	if err != nil {
		return responses.ProductErrorResponse(c, 400, err.Error())
	}

	return responses.ProductSuccessResponse(c, 200, message, nil)
}

func (pc *ProductController) DeleteProduct(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return responses.ProductErrorResponse(c, 400, invalidProductID)
	}

	message, err := pc.ProductRepo.DeleteProduct(c.Request().Context(), id)
	if err != nil {
		return responses.ProductErrorResponse(c, 400, err.Error())
	}

	return responses.ProductSuccessResponse(c, 200, message, nil)
}
