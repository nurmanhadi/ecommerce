package middleware

import (
	"ecommerce/pkg/infrastructure/exception"
	"ecommerce/pkg/infrastructure/response"
	"ecommerce/pkg/infrastructure/security"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token, err := getTokenFromHeader(c)
	if err != nil {
		return response.ResponseError(c, err)
	}
	userId, err := security.VerifyJwt(token)
	if err != nil {
		return response.ErrorResponse(c, 401, err.Error())
	}
	c.Locals("userId", userId)
	return c.Next()
}

func getTokenFromHeader(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return "", exception.AuthTokenIsNull
	}
	token := strings.Split(authHeader, " ")
	if token[0] != "Bearer" {
		return "", exception.AuthBearerReqired
	}
	return token[1], nil
}
