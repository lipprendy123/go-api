package main

import (
	"fmt"
	"log"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	 dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
  	 _, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection error")
	}

	fmt.Println("Db connection successfully")
	
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)

	v1.GET("/hello", handler.HelloHandler)

	v1.GET("/book/:id", handler.BookHandler)

	v1.GET("/query", handler.QueryHandler)

	v1.POST("/books", handler.PostBooksHandler)
	
	router.Run()
}
