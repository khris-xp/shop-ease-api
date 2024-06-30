package controllers

import (
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/khris-xp/shop-ease-api/configs"
	"github.com/khris-xp/shop-ease-api/models"
	"github.com/khris-xp/shop-ease-api/repositories"
	"github.com/khris-xp/shop-ease-api/responses"
	"github.com/labstack/echo/v4"
)

var (
	jwtSecret = []byte(configs.EnvSecretKey())
)

type AuthController struct {
	UserRepo *repositories.UserRepository
}

func NewAuthController(userRepo *repositories.UserRepository) *AuthController {
	return &AuthController{UserRepo: userRepo}
}

func (ac *AuthController) RegisterUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return responses.UserErrorResponse(c, 400, "invalid request body")
	}

	tokenString, err := ac.UserRepo.RegisterUser(c.Request().Context(), user)
	if err != nil {
		return responses.UserErrorResponse(c, 400, "email already exists")
	} else if tokenString == "" {
		return responses.UserErrorResponse(c, 400, "unable to register user")
	}

	return responses.AuthUserSuccessResponse(c, 200, "success", tokenString)
}

func (ac *AuthController) LoginUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return responses.UserErrorResponse(c, 400, "invalid request body")
	}

	tokenString, err := ac.UserRepo.LoginUser(c.Request().Context(), user.Email, user.Password)
	if err != nil {
		return responses.UserErrorResponse(c, 400, "invalid email or password")
	} else if tokenString == "" {
		return responses.UserErrorResponse(c, 400, "unable to login user")
	}

	return responses.AuthUserSuccessResponse(c, 200, "success", tokenString)
}

func (ac *AuthController) GetUser(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(401, map[string]string{"message": "Unauthorized"})
	}

	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.JSON(401, map[string]string{"message": "Unauthorized"})
	}

	token = parts[1]

	if token == "" {
		return responses.UserErrorResponse(c, 401, "unauthorized")
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !parsedToken.Valid {
		return responses.UserErrorResponse(c, 401, "unauthorized")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return responses.UserErrorResponse(c, 401, "unauthorized")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return responses.UserErrorResponse(c, 401, "unauthorized")
	}

	user, err := ac.UserRepo.GetUserProfile(c.Request().Context(), email)
	if err != nil {
		return responses.UserErrorResponse(c, 400, "unable to get user")
	}

	return responses.UserSuccessResponse(c, 200, "success", user)
}

func (ac *AuthController) RefreshToken(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(401, map[string]string{"message": "Unauthorized"})
	}

	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.JSON(401, map[string]string{"message": "Unauthorized"})
	}

	token = parts[1]

	if token == "" {
		return responses.UserErrorResponse(c, 401, "unauthorized")
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !parsedToken.Valid {
		return responses.UserErrorResponse(c, 401, "unauthorized")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return responses.UserErrorResponse(c, 401, "unauthorized")
	}

	_, ok = claims["email"].(string)
	if !ok {
		return responses.UserErrorResponse(c, 401, "unauthorized")
	}

	tokenString, err := ac.UserRepo.RefreshToken(c.Request().Context(), token)

	if err != nil {
		return responses.UserErrorResponse(c, 400, "unable to refresh token")
	}

	return responses.AuthUserSuccessResponse(c, 200, "success", tokenString)
}
