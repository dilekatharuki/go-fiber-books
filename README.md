Here's your updated and detailed `README.md` file including **pagination**, **search**, setup instructions, and test coverage.

---

```markdown
# 📚 Go Fiber Books API

A simple RESTful API built with [Go Fiber](https://gofiber.io/) and [GORM](https://gorm.io/) that allows you to manage a collection of books with full CRUD functionality, pagination, and search.

---

## 🚀 Features

- ✅ Create, Read, Update, Delete books
- 🔍 Search by `title` or `author`
- 📄 Paginated book listing
- 💾 SQLite database integration
- ✅ Unit tests for all CRUD operations
- ⚙️ Modular folder structure with services and handlers

---
## 📦 Tech Stack

- **Go** - Programming language
- **Fiber** - Web framework
- **GORM** - ORM for SQLite
- **SQLite** - Lightweight database

---

## 🔧 Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/dilekatharuki/go-fiber-books.git
cd go-fiber-books
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the server

```bash
go run main.go
```

Server starts at: `http://localhost:3000`

---

## 🧪 Run Tests

```bash
go test -v ./handlers -count=1
```

- `-v` = Verbose
- `-count=1` = Ensures tests aren't cached

---

## 📚 API Endpoints

### ➕ Create a Book

```
POST /books
```

**Request Body:**

```json
{
  "title": "Clean Code",
  "author": "Robert C. Martin",
  "year": 2008
}
```

---

### 📖 Get All Books (with Pagination + Search)

```
GET /books?page=1&limit=10&search=clean
```

| Query Param | Description                          |
|-------------|--------------------------------------|
| `page`      | Page number (default: 1)             |
| `limit`     | Number of results per page (default: 10) |
| `search`    | Filter books by title or author      |

---

### 🔍 Get a Single Book

```
GET /books/:id
```

---

### ✏️ Update a Book

```
PUT /books/:id
```

**Request Body:**

```json
{
  "title": "Updated Title",
  "author": "Updated Author",
  "year": 2023
}
```

---

### ❌ Delete a Book

```
DELETE /books/:id
```

---

## 🔍 Example Usage

### Search for "Golang" books, 5 per page:

```
GET /books?page=1&limit=5&search=golang
```

## 👨‍💻 Author

Developed by [Dileka Tharuki ]  
For internship/academic project use.
```
