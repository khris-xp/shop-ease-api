package middlewares

import (
	"github.com/khris-xp/shop-ease-api/configs"
	"github.com/labstack/echo/v4"
)

var (
	jwtSecret = []byte(configs.EnvSecretKey())
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString, err := getTokenString(c)
		if err != nil {
			return c.JSON(401, map[string]string{"message": "Unauthorized"})
		}

		token, err := parseToken(tokenString)
		if err != nil {
			return c.JSON(401, map[string]string{"message": "Unauthorized"})
		}

		if !token.Valid {
			return c.JSON(401, map[string]string{"message": "Unauthorized"})
		}

		return next(c)
	}
}
