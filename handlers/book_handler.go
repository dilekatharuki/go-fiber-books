package handlers

import (
	"strconv"

	"github.com/dilekatharuki/go-fiber-books/models"
	"github.com/dilekatharuki/go-fiber-books/services"
	"github.com/gofiber/fiber/v2"
)

// CreateBook handles POST /books
func CreateBook(c *fiber.Ctx) error {
    book := new(models.Book)
    if err := c.BodyParser(book); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    if err := services.CreateBook(book); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(fiber.StatusCreated).JSON(book)
}

// GetBooks handles GET /books
func GetBooks(c *fiber.Ctx) error {
    page, err := strconv.Atoi(c.Query("page", "1"))
    if err != nil || page < 1 {
        page = 1
    }

    limit, err := strconv.Atoi(c.Query("limit", "10"))
    if err != nil || limit < 1 {
        limit = 10
    }

    search := c.Query("search", "")

    books, err := services.GetBooks(page, limit, search)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch books",
        })
    }

    return c.JSON(books)
}


// GetBook handles GET /books/:id
func GetBook(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    book, err := services.GetBook(uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
    }
    return c.JSON(book)
}

// UpdateBook handles PUT /books/:id
func UpdateBook(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    book := new(models.Book)
    if err := c.BodyParser(book); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    updatedBook, err := services.UpdateBook(uint(id), book)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(updatedBook)
}

// DeleteBook handles DELETE /books/:id
func DeleteBook(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    if err := services.DeleteBook(uint(id)); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.SendStatus(fiber.StatusNoContent)
}
