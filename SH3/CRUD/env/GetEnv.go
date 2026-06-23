package env

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetEnv(c *fiber.Ctx) error {
	if value, exists := os.LookupEnv("SECRET"); exists {
		return c.JSON(fiber.Map{
			"SECRET": value,
		})
	}
	return c.JSON(fiber.Map{
		"SECRET": "defaultsecret",
	})
}

func GetEnvFile(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"SECRET": os.Getenv("SECRET_KEY"),
	})
}
