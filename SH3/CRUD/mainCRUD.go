package crud

import (
	"SH3/CRUD/env"
	middleware "SH3/Middleware"
	"SH3/views"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func CRUD() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env Error")
	}
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	books = append(books, Book{ID: 1, Title: "2000", Author: "George Orwell"})
	books = append(books, Book{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})

	app.Post("/login", middleware.Login)
	app.Get("/test-html", views.TestHTML)
	app.Get("/config", env.GetEnv)
	app.Get("/configFile", env.GetEnvFile)
	app.Post("/uploda", UploadFile)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	app.Use(middleware.CheckMiddleware)

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)

	app.Post("/books", CreateBook)

	app.Put("/books/:id", UpdateBook)

	app.Delete("/books/:id", DeleteBook)

	app.Listen(":8080")
}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}
