package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"pustaka-api/book"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
			"name": "Alif Rendy",
			"bio":  "Just ordinary people",
		})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
			"title" : "Hello World!!",
			"subtitle" : "Belajar golang",
		})
}

func BookHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

func PostBooksHandler(c *gin.Context) {
	
	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {
    var errorMessages []string 

    if validationErrors, ok := err.(validator.ValidationErrors); ok {
        for _, e := range validationErrors {
            errorMessage := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
            errorMessages = append(errorMessages, errorMessage) 
        }

        c.JSON(http.StatusBadRequest, gin.H{
            "errors": errorMessages,
        })
        return
    }

    c.JSON(http.StatusInternalServerError, gin.H{
        "error": err.Error(),
    })
    return
}


	c.JSON(http.StatusOK, gin.H {
		"title" : bookInput.Title,
		"price" : bookInput.Price,
	}) 
}