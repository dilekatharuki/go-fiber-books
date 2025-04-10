package main

import (
	"log"

	"github.com/dilekatharuki/go-fiber-books/database"
	"github.com/dilekatharuki/go-fiber-books/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
    // Initialize the database
    database.Connect()

    // Create a new Fiber app
    app := fiber.New()

    // Routes
    app.Post("/books", handlers.CreateBook)
    app.Get("/books", handlers.GetBooks)
    app.Get("/books/:id", handlers.GetBook)
    app.Put("/books/:id", handlers.UpdateBook)
    app.Delete("/books/:id", handlers.DeleteBook)

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}
