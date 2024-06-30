package responses

import (
	"github.com/khris-xp/shop-ease-api/types"
	"github.com/labstack/echo/v4"
)

func AuthUserSuccessResponse(c echo.Context, statusCode int, message string, token string) error {
	return c.JSON(statusCode, types.AuthResponse{
		Status:  statusCode,
		Message: message,
		Token:   token,
	})
}

func UserSuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, types.UserResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	})
}

func UserErrorResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, types.UserResponse{
		Status:  statusCode,
		Message: message,
		Data:    nil,
	})
}