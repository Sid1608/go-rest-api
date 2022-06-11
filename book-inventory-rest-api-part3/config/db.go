package db
import (

	"fmt"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myapp/model"
	_ "github.com/go-sql-driver/mysql"
	
)
var DB *gorm.DB
var err error
const DNS="root:siddharth@tcp(127.0.0.1:3306)/booksdb?charset=utf8&parseTime=True&loc=Local"
func InitialMigration(){
	DB,err=gorm.Open(mysql.Open(DNS),&gorm.Config{})
	if err!=nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	fmt.Println("Connecting to Db")
	DB.AutoMigrate(&model.Book{})

}

func GetDBInstance()*gorm.DB{
	return DB
}
