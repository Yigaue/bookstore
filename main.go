package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yigaue/bookstore/database"
	"github.com/yigaue/bookstore/models"
)

// books slice to seed record to book store.
var books = []models.Book{
	{ID: "1", Title: "A Day in the Life of Abed Salama", Author: "Nathan Thrall", Price: 56.99},
	{ID: "2", Title: "King: A life", Author: "Jonathan Eig", Price: 56.99},
	{ID: "3", Title: "Where we go from here", Author: "Bernie Sanders", Price: 17.99},
	{ID: "4", Title: "Buiding a dream server", Author: "Yiga ue", Price: 39.99},
	{ID: "5", Title: "Clean Code ", Author: "Robert C Martin", Price: 39.99},
}

// getBooks responds with the list of all books as json
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	database.DBConnect()
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
