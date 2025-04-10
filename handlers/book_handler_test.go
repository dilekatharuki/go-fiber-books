package handlers

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/dilekatharuki/go-fiber-books/database"
	"github.com/dilekatharuki/go-fiber-books/models"
	"github.com/gofiber/fiber/v2"
)

func setupApp() *fiber.App {
	database.Connect()
	app := fiber.New()

	app.Post("/books", CreateBook)
	app.Get("/books", GetBooks)
	app.Get("/books/:id", GetBook)
	app.Put("/books/:id", UpdateBook)
	app.Delete("/books/:id", DeleteBook)

	return app
}

func createTestBook(app *fiber.App) uint {
	req := httptest.NewRequest("POST", "/books", strings.NewReader(`{"title":"Test Book", "author":"Test Author", "year":2020}`))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	body, _ := io.ReadAll(resp.Body)
	var book models.Book
	json.Unmarshal(body, &book)
	return book.ID
}

func TestCreateBook(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("POST", "/books", strings.NewReader(`{"title":"New Book", "author":"Author X", "year":2021}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	if resp.StatusCode != fiber.StatusCreated {
		t.Errorf("Expected status %d, got %d", fiber.StatusCreated, resp.StatusCode)
	}
}

func TestGetBooks(t *testing.T) {
	app := setupApp()
	createTestBook(app)

	req := httptest.NewRequest("GET", "/books", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("GET /books error: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status %d, got %d", fiber.StatusOK, resp.StatusCode)
	}
}

func TestGetBook(t *testing.T) {
	app := setupApp()
	id := createTestBook(app)

	req := httptest.NewRequest("GET", "/books/"+strconv.Itoa(int(id)), nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("GET /books/:id error: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status %d, got %d", fiber.StatusOK, resp.StatusCode)
	}
}

func TestUpdateBook(t *testing.T) {
	app := setupApp()
	id := createTestBook(app)

	updateBody := `{"title":"Updated Book", "author":"Updated Author", "year":2023}`
	req := httptest.NewRequest("PUT", "/books/"+strconv.Itoa(int(id)), strings.NewReader(updateBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("PUT /books/:id error: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status %d, got %d", fiber.StatusOK, resp.StatusCode)
	}
}

func TestDeleteBook(t *testing.T) {
	app := setupApp()
	id := createTestBook(app)

	req := httptest.NewRequest("DELETE", "/books/"+strconv.Itoa(int(id)), nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("DELETE /books/:id error: %v", err)
	}
	if resp.StatusCode != fiber.StatusNoContent {
		t.Errorf("Expected status %d, got %d", fiber.StatusNoContent, resp.StatusCode)
	}
}
