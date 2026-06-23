package crud

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UpdateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookupdate := new(Book)
	if err := c.BodyParser(bookupdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books[i].Title = bookupdate.Title
			books[i].Author = bookupdate.Author
			return c.JSON(books[i])
		}

	}
	return c.SendStatus(fiber.StatusNotFound)
}
