package response

import "github.com/gofiber/fiber/v2"

func ResponseSuccess(c *fiber.Ctx, statusCode int, message string, data map[string]interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "success",
		"message": message,
		"data":    data,
		"links": map[string]string{
			"self": c.OriginalURL(),
		},
	})
}
