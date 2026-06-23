package httpserver

import (
	"github.com/gofiber/fiber/v2"
)

func ShowHelloFiber() {
	app := fiber.New()

	app.Get("/hellofiber", func(c *fiber.Ctx) error {
		return c.SendString("Heool Fiber!")
	})

	app.Listen(":8080")
}
