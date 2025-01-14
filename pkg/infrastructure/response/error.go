package response

import (
	"ecommerce/pkg/infrastructure/exception"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(c *fiber.Ctx, statusCode int, message string, err string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "error",
		"message": message,
		"errors":  err,
		"links": map[string]string{
			"self": c.OriginalURL(),
		},
	})
}
func ResponseError(c *fiber.Ctx, err error) error {
	if err != nil {
		if errors.Is(err, exception.JwtExpired) {
			return ErrorResponse(c, 401, "unauthorized", err.Error())
		} else if errors.Is(err, exception.JwtInvalidSignature) {
			return ErrorResponse(c, 401, "unauthorized", err.Error())
		} else if errors.Is(err, exception.UserEmailAlreadyExist) {
			return ErrorResponse(c, 400, "baad request", err.Error())
		} else if errors.Is(err, exception.UserEmailAndPasswordIsWrong) {
			return ErrorResponse(c, 400, "baad request", err.Error())
		} else if finalErr, ok := err.(*exception.ValidationError); ok {
			return ErrorResponse(c, 400, "baad request", finalErr.Message)
		} else {
			return ErrorResponse(c, 500, "internal server error", err.Error())
		}
	}
	return nil
}
