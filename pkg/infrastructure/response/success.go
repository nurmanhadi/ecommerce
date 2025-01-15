package response

import "github.com/gofiber/fiber/v2"

func ResponseSuccess(c *fiber.Ctx, statusCode int, data any) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"data": data,
		"links": map[string]string{
			"self": c.OriginalURL(),
		},
	})
}
