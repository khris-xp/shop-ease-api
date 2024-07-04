package controllers

import (
	"github.com/khris-xp/shop-ease-api/models"
	"github.com/khris-xp/shop-ease-api/repositories"
	"github.com/khris-xp/shop-ease-api/responses"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryController struct {
	CategoryRepo *repositories.CategoryRepository
}

func NewCategoryController(categoryRepo *repositories.CategoryRepository) *CategoryController {
	return &CategoryController{CategoryRepo: categoryRepo}
}

func (cc *CategoryController) CreateCategory(c echo.Context) error {
	var category models.Category
	if err := c.Bind(&category); err != nil {
		return responses.CategoryErrorResponse(c, 400, "invalid request body")
	}

	message, err := cc.CategoryRepo.CreateCategory(c.Request().Context(), category)
	if err != nil {
		return responses.CategoryErrorResponse(c, 400, err.Error())
	}

	return responses.CategorySuccessResponse(c, 201, message, nil)
}

func (cc *CategoryController) GetCategories(c echo.Context) error {
	categories, err := cc.CategoryRepo.GetCategories(c.Request().Context())
	if err != nil {
		return responses.CategoryErrorResponse(c, 400, err.Error())
	}

	return responses.CategorySuccessResponse(c, 200, "success", categories)
}

const invalidCategoryID = "invalid category ID"

func (cc *CategoryController) GetCategoryByID(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return responses.CategoryErrorResponse(c, 400, invalidCategoryID)
	}

	category, err := cc.CategoryRepo.GetCategoryByID(c.Request().Context(), id)
	if err != nil {
		return responses.CategoryErrorResponse(c, 400, err.Error())
	}

	return responses.CategorySuccessResponse(c, 200, "success", category)
}

func (cc *CategoryController) UpdateCategory(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return responses.CategoryErrorResponse(c, 400, "invalid category ID")
	}

	var category models.Category
	if err := c.Bind(&category); err != nil {
		return responses.CategoryErrorResponse(c, 400, "invalid request body")
	}

	message, err := cc.CategoryRepo.UpdateCategory(c.Request().Context(), id, category)
	if err != nil {
		return responses.CategoryErrorResponse(c, 400, err.Error())
	}

	return responses.CategorySuccessResponse(c, 200, "Update Category Sucess", message)
}

func (cc *CategoryController) DeleteCategory(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return responses.CategoryErrorResponse(c, 400, "invalid category ID")
	}

	message, err := cc.CategoryRepo.DeleteCategory(c.Request().Context(), id)
	if err != nil {
		return responses.CategoryErrorResponse(c, 400, err.Error())
	}

	return responses.CategorySuccessResponse(c, 200, "Delete Category Success", message)
}
