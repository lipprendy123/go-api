package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/hello", helloHandler)

	router.GET("/book/:id", bookHandler)

	router.GET("/query", queryHandler)

	router.POST("/books", postBooksHandler)
	
	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
			"name": "Alif Rendy",
			"bio":  "Just ordinary people",
		})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
			"title" : "Hello World!!",
			"subtitle" : "Belajar golang",
		})
}

func bookHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

type BookInput struct {
	Title string
	Price int
	SubTitle string `json:"sub_title"`
}

func postBooksHandler(c *gin.Context) {
	
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H {
		"title" : bookInput.Title,
		"price" : bookInput.Price,
		"sub_title" : bookInput.SubTitle,
	}) 
}

