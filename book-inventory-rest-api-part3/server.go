package main

import (
	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"myapp/controller"
	"myapp/config"
	
	// _ "github.com/go-sql-driver/mysql"
	
)



//----------


func main() {
	db.InitialMigration()
	e := echo.New()
	// DB:=db.GetDBInstance()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.GET("/",controller.CheckServer)
	e.GET("/books", controller.GetAllBooks)
	e.POST("/createBook", controller.CreateBook)
	e.GET("/books/:id", controller.GetBook)
	e.GET("/bookscount", controller.GetBooksCount)
	e.GET("/authors", controller.GetAuthors)
	e.GET("authorsBook/:author",controller.GetBooksByAuthor)
	// e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}