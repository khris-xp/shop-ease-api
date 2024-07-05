package responses

import (
	"github.com/khris-xp/shop-ease-api/types"
	"github.com/labstack/echo/v4"
)

func ProductSuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, types.ProductResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	})
}

func ProductErrorResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, types.ProductResponse{
		Status:  statusCode,
		Message: message,
		Data:    nil,
	})
}
