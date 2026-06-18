package httpserver

import (
	"github.com/gofiber/fiber/v2"
)

func helloHandler() {

}

func EchoData() {
	app := fiber.New()

	app.Get("/hello",
		func(c *fiber.Ctx) error {
			return c.SendString("Hello Word! test_1")
		})

	app.Listen(":8080")
}
