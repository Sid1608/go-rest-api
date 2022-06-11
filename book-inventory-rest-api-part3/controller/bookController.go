package controller


import (
	"net/http"
	"encoding/json"
	"strconv"
	"fmt"
	"myapp/model"
	"myapp/config"
	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)
// var DB *gorm.DB=db.GetDBInstance()
// type ExportedType interface {
//     getAllBooks() string

func CheckServer(c echo.Context) error {
	fmt.Println("server is running successfully on port 1323")
	book:="Hello"
	return c.JSON(http.StatusOK,book)
}
func CreateBook(c echo.Context) error {
	fmt.Println("Creating Book",c)
	var book model.Book
	var DB *gorm.DB=db.GetDBInstance()
	err := json.NewDecoder(c.Request().Body).Decode(&book)
     if err != nil {

         log.Error("empty json body")
         return nil
     }
	 DB.Create(&book)
	return c.JSON(http.StatusOK,book)
}

func GetBook(c echo.Context) error {
	var DB *gorm.DB=db.GetDBInstance()
	id, _ := strconv.Atoi(c.Param("id"))
	var book model.Book
	DB.First(&book,id)
	return c.JSON(http.StatusOK, book)
	// defer DB.Close()
}
func GetBooksCount(c echo.Context) error{
	var DB *gorm.DB=db.GetDBInstance()
	var books int64
	DB.Model(&model.Book{}).Distinct("title").Count(&books)
	// defer books.Close()
	return c.JSON(http.StatusOK, books)
}



func GetAllBooks(c echo.Context) error {
	var DB *gorm.DB=db.GetDBInstance()
	var books []model.Book
	DB.Find(&books)
	// defer books.Close()
	return c.JSON(http.StatusOK, books)
}
func GetAuthors(c echo.Context) error {
	var DB *gorm.DB=db.GetDBInstance()
	var books []model.Book
	DB.Distinct("author").Find(&books)

	// DB.Find(&books)
	var authors []string 
	for _, book := range books{
		authors = append(authors, book.Author)
	}
	// defer books.Close()
	return c.JSON(http.StatusOK, authors)
}
func GetBooksByAuthor(c echo.Context)error{
	var DB *gorm.DB=db.GetDBInstance()
	author:=c.Param("author")
	var books []model.Book
	DB.Where("Author = ?", author).Find(&books)
	// defer books.Close()
	return c.JSON(http.StatusOK, books)
}
