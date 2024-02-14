package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yigaue/bookstore/database"
	"github.com/yigaue/bookstore/models"
)

// getBooks responds with the list of all books as json
func getBooks(c *gin.Context) {
	var db *sql.DB
	db = database.DBConnect()
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

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.POST("/books", postBooks)
	router.Run("localhost:8080")
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func postBooks(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
