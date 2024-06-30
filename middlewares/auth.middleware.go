package middlewares

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/khris-xp/shop-ease-api/configs"
	"github.com/labstack/echo/v4"
)

var (
	jwtSecret = []byte(configs.EnvSecretKey())
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(401, map[string]string{"message": "Unauthorized"})
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(401, map[string]string{"message": "Unauthorized"})
		}

		tokenString = parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, nil
			}
			return jwtSecret, nil
		})

		if err != nil {
			return c.JSON(401, map[string]string{"message": "Unauthorized"})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user", claims["user"])
			return next(c)
		}

		return c.JSON(401, map[string]string{"message": "Unauthorized"})
	}
}
