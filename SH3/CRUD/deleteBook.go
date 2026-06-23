package crud

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}

	}
	return c.SendStatus(fiber.StatusNotFound)
}
