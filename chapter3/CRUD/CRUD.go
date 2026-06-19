package crud

import (
	env "learngo/chapter3/ENV"
	"learngo/chapter3/uplodafile"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func ShowdataCRUD() {
	engine := html.New("./chapter3/views", ".html")
	appcrud := fiber.New(fiber.Config{
		Views: engine,
	})

	books = append(books, Book{ID: 1, Title: "1984", Author: "George Orwell"})
	books = append(books, Book{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})

	appcrud.Get("/books", getAllBooks)
	appcrud.Get("/books/:id", getAllBook)
	appcrud.Post("/books", CreateBook)
	appcrud.Put("/books/:id", UpdateBook)
	appcrud.Delete("/books/:id", DeleteBook)

	appcrud.Post("/upload", uplodafile.UploadFile)
	appcrud.Get("/test-html", TestHTML)

	appcrud.Get("/config", env.GetEnv)

	appcrud.Listen(":8080")

}

func getAllBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

// ค้นหาชุดข้อมูลและดัก err
func getAllBook(c *fiber.Ctx) error {
	bookid, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if book.ID == bookid {
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

// เพิ่มข้อมูล ถ้า run ใหม่ข้อมูลจะหาย
// func createBook(c *fiber.Ctx) error {

// 	book := new(Book)
// 	if err := c.BodyParser(book); err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}
// 	books = append(books, *book)
// 	return c.JSON(book)
// }

// update ข้อมูล ถ้า run ใหม่ข้อมูลจะหาย
// func updateBook(c *fiber.Ctx) error {
// 	bookid, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}

// 	bookUpdate := new(Book)
// 	if err := c.BodyParser(bookUpdate); err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}

// 	for i, book := range books {
// 		if book.ID == bookid {
// 			books[i].Title = bookUpdate.Title
// 			books[i].Author = bookUpdate.Author
// 			return c.JSON(books[i])
// 		}
// 	}

// 	return c.SendStatus(fiber.StatusNotFound)
// }

// func deleteBook(c *fiber.Ctx) error {
// 	bookid, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}

// 	for i, book := range books {
// 		if book.ID == bookid {
// 			books = append(books[:i], books[i+1:]...)
// 			return c.SendStatus(fiber.StatusNoContent)
// 		}
// 	}
// 	return c.SendStatus(fiber.StatusNotFound)
// }
