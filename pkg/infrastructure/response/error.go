package response

import (
	"ecommerce/pkg/infrastructure/exception"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(c *fiber.Ctx, statusCode int, err string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status": "error",
		"errors": err,
		"links": map[string]string{
			"self": c.OriginalURL(),
		},
	})
}
func ResponseError(c *fiber.Ctx, err error) error {
	if err != nil {
		if errors.Is(err, exception.JwtExpired) {
			return ErrorResponse(c, 401, err.Error())
		} else if errors.Is(err, exception.AuthBearerReqired) {
			return ErrorResponse(c, 401, err.Error())
		} else if errors.Is(err, exception.AuthTokenIsNull) {
			return ErrorResponse(c, 401, err.Error())
		} else if errors.Is(err, exception.JwtInvalidSignature) {
			return ErrorResponse(c, 401, err.Error())
		} else if errors.Is(err, exception.UserEmailAlreadyExist) {
			return ErrorResponse(c, 400, err.Error())
		} else if errors.Is(err, exception.OrderStatusQueryIsRequired) {
			return ErrorResponse(c, 400, err.Error())
		} else if errors.Is(err, exception.UserEmailAndPasswordIsWrong) {
			return ErrorResponse(c, 400, err.Error())
		} else if validationErr, ok := err.(validator.ValidationErrors); ok {
			var values []string
			for _, fieldErr := range validationErr {
				value := fmt.Sprintf("field %s is %s %s", fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
				values = append(values, value)
			}
			str := strings.Join(values, ", ")
			return ErrorResponse(c, 400, str)
		} else if errors.Is(err, exception.ProductNotFound) {
			return ErrorResponse(c, 404, err.Error())
		} else if errors.Is(err, exception.UserNotFound) {
			return ErrorResponse(c, 404, err.Error())
		} else if errors.Is(err, exception.OrderNotFound) {
			return ErrorResponse(c, 404, err.Error())
		} else {
			return ErrorResponse(c, 500, "internal server error")
		}
	}
	return nil
}
