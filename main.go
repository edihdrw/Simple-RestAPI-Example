package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/book_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection error!")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/book", bookHandler.GetBooks)
	v1.GET("/book/:id", bookHandler.GetBook)
	v1.POST("/book", bookHandler.CreateBook)
	v1.PUT("/book/:id", bookHandler.UpdateBook)
	v1.DELETE("/book/:id", bookHandler.DeleteBook)

	router.Run()

}
