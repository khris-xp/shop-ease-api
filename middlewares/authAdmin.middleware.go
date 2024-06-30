package middlewares

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func AuthAdminMiddlewares(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString, err := getTokenString(c)
		if err != nil {
			return c.JSON(401, map[string]string{"message": err.Error()})
		}

		token, err := parseToken(tokenString)
		if err != nil {
			return c.JSON(401, map[string]string{"message": "Unauthorized"})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if !isAdmin(claims) {
				return c.JSON(401, map[string]string{"message": "Unauthorized Admin Access"})
			}
			c.Set("user", claims["user"])
			return next(c)
		}

		return c.JSON(401, map[string]string{"message": "Unauthorized"})
	}
}

func getTokenString(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Unauthorized")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("Unauthorized")
	}

	return parts[1], nil
}

func parseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return jwtSecret, nil
	})
}

func isAdmin(claims jwt.MapClaims) bool {
	return claims["role"] == "admin"
}
