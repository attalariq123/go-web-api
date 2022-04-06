package main

import (
	"web-api/database"
	"web-api/handler"
	"web-api/repository"
	"web-api/service"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.InitDB()
	database.MigrateDB()

	bookRepository := repository.NewRepo(db)
	bookService := service.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/book", bookHandler.GetBooksHandler)
	v1.POST("/book", bookHandler.PostBookHandler)
	v1.GET("book/:id", bookHandler.GetBookById)
	v1.PUT("book/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("book/:id", bookHandler.DeleteBookHandler)
	v1.GET("/user", userHandler.GetUsersHandler)
	v1.POST("/user", userHandler.PostUserHandler)
	v1.GET("/user/:email", userHandler.FindByEmail)

	router.Run()

	/**
	WORKFLOW
	main
	handler
	service
	repository
	db
	mysql
	**/
}
