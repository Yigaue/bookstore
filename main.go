package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yigaue/bookstore/database"
	"github.com/yigaue/bookstore/models"
)

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.POST("/books", postBooks)
	router.DELETE("/books/:id", deleteBook)
	router.Run("localhost:8080")
}

// getBooks responds with the list of all books as json
func getBooks(c *gin.Context) {
	// var db *sql.DB
	db := database.DBConnect()
	var books []models.Book

	rows, err := db.Query("SELECT * FROM book")
	if err != nil {
		fmt.Errorf("getBooks: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Price); err != nil {
			fmt.Errorf("getBooks: %v", err)
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		fmt.Errorf("getBooks: %v", err)
	}
	c.IndentedJSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
	db := database.DBConnect()
	id := c.Param("id")
	var book models.Book

	row := db.QueryRow("SELECT * FROM book WHERE id = ?", id)
	err := row.Scan(&book.ID, &book.Author, &book.Title, &book.Price)

	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
  }

	if err != nil {
		fmt.Errorf("book ID, %d: %v", id, err)
	}

	c.IndentedJSON(http.StatusOK, book)
}

func postBooks(c *gin.Context) {
	db := database.DBConnect()
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	row, err := db.Exec("INSERT INTO book (title, author, price) VALUES (?, ?, ?)", newBook.Title, newBook.Author, newBook.Price)
	if err != nil {
		fmt.Errorf("postBooks %v", err)
	}

	id, err := row.LastInsertId()

	if err != nil {
		fmt.Errorf("error getting lastID: %v", err)
	}

	newBook.ID = strconv.Itoa(int(id))

	c.IndentedJSON(http.StatusCreated, newBook)
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	db := database.DBConnect()
	_, err := db.Exec("DELETE FROM book WHERE id = ?", id)

	if err != nil {
		fmt.Errorf("deleteBooks %v", err)
	}

	if err != nil {
		fmt.Errorf("Error getting lastInsertedId deleteBook %v", err)
	}

	c.IndentedJSON(http.StatusOK, "Book deleted successful")
}
