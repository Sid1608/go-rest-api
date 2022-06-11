package main
import (
  "github.com/gorilla/mux"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "net/http"
  "encoding/json"
  "fmt"
  "io/ioutil"
)
type Book struct{
	
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}
var db *sql.DB
var err error
func getBooks(w http.ResponseWriter, r *http.Request) {
	
	var Books []Book
	result, err := db.Query("SELECT id, title,author from books")
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

	  var id string
	  var title string 
	  var author string 
	  err := result.Scan(&id,&title,&author)
	
	   book:=Book{id,title,author}
	  if err != nil {
		panic(err.Error())
	  }
	  
	  Books = append(Books, book)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)
  }
  func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO books(title,author) VALUES(?,?)")
	if err != nil {
	  panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	  panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]
	author:=keyVal["author"]
	_, err = stmt.Exec(title,author)
	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "New Book was created")
  }
  func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT id,title, author FROM books WHERE id = ?", params["id"])
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()
	var book Book
	for result.Next() {
	  err := result.Scan(&book.ID,&book.Title,&book.Author)
	  if err != nil {
		panic(err.Error())
	  }
	}
	json.NewEncoder (w).Encode(book)
  }
  func getBooksCount(w http.ResponseWriter, r *http.Request){
	// w.Header().Set("Content-Type", "application/json")
	var Books []Book
	result, err := db.Query("SELECT id, title, author from books")
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
	  var book Book
	  err := result.Scan(&book.ID,&book.Title,&book.Author)
	  if err != nil {
		panic(err.Error())
	  }
	  Books = append(Books, book)
	}
	res,_:=json.Marshal(len(Books))
	  w.Header().Set("Content-Type", "application/json")
	  w.WriteHeader(http.StatusOK)
	  w.Write(res)
  }
  func getAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Books []string
	result, err := db.Query("SELECT id, title, author from books")
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
	  var book Book
	  err := result.Scan(&book.ID,&book.Title,&book.Author)
	  if err != nil {
		panic(err.Error())
	  }
	  Books = append(Books, book.Author)
	}
	json.NewEncoder(w).Encode(Books)
  }
  func getBooksByAuthur(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var Books []Book
	result, err := db.Query("SELECT id, title, author from books")
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
	  var book Book
	  err := result.Scan(&book.ID,&book.Title,&book.Author)
	  if err != nil {
		panic(err.Error())
	  }
	  if book.Author == params["authname"] {
	  	Books = append(Books, book)
	  }
	}
	json.NewEncoder(w).Encode(Books)
  }
func main() {
	db, err = sql.Open("mysql", "root:siddharth@tcp(127.0.0.1:3306)/bookstore")
	if err != nil {
		panic(err.Error())
	}
  defer db.Close()
  router := mux.NewRouter()
  router.HandleFunc("/books", getBooks).Methods("GET")
  router.HandleFunc("/storeBook", createBook).Methods("POST")
  router.HandleFunc("/books/{id}", getBook).Methods("GET")
  router.HandleFunc("/bookscount", getBooksCount).Methods("GET")
  router.HandleFunc("/bookAuthors", getAuthors).Methods("GET")
  router.HandleFunc("/authorbooks/{authname}", getBooksByAuthur).Methods("GET")
  http.ListenAndServe(":8000", router)
}

