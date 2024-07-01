package responses

import (
	"github.com/khris-xp/shop-ease-api/types"
	"github.com/labstack/echo/v4"
)

func CategorySuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, types.CategoryResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	})
}

func CategoryErrorResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, types.CategoryResponse{
		Status:  statusCode,
		Message: message,
		Data:    nil,
	})
}
