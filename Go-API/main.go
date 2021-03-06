package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct{
	ID  string   `json: "id"`   //upper case makes it an exported field, but when serialising in json convert the field name to lower case
	Title string  `json: "title"`
	Author string  `json: "author"`
	Quantity int	`json: "quantity"`
}

//mock data is this slice of books (will be a database in a real life example)
var books = []book{
				{ID: "1",  Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
				{ID: "2",  Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
				{ID: "3",  Title: "War And Peace", Author: "Leo Tolstoy", Quantity: 6},
}





//set up router to handle main end points of api by implementing the book handling functions first:


//getting all the books
func getBooks(c*gin.Context){
		c.IndentedJSON(http.StatusOK, books)   //nicely indented json, the status that we are sending is status okay, the data that we are sending is "books". So we return a json obejct that has all of the books in it
}


//returning a specic book
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."}) //return custom request for bad request or book not found
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}


//getting a book by id
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

//* symbol is used to declare a pointer and to dereference, as well as changing the value of the pointer location as well  & symbol points to the address of the stored value.

//checking out a specific book
func checkoutBook(c *gin.Context){
	id, ok := c.GetQuery("id") //chech query parameter in the path built in function in gin framework

	if !ok{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	book, err := getBookById(id)

	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."}) //return custom request for bad request or book not found
		return
	}

	if book.Quantity <= 0{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."}) //return custom request for bad request or book not found
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}


func returnBook(c *gin.Context){

	id, ok := c.GetQuery("id") //chech query parameter in the path built in function in gin framework

	if !ok{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}




//creating a new book
func createBook(c*gin.Context){ //c stores query parameters, headers

	var newBook book   	// the new book is of type book

	if err := c.BindJSON(&newBook); err != nil{

			// if the error is not equal to null, in that case we shall simpy return
			return
		}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}






func main(){

	//create a gin router
	//can route a specific route to a function with router variable
	router := gin.Default()


	//return all books
	router.GET("/books", getBooks)   //when visting the /books endpoint the getbooks function is called


	//return a specific book
	router.GET("/books/:id", bookById)


	//create a book:
	router.POST("/books", createBook)


	//checkout a book
	router.PATCH("/checkout", checkoutBook)

	//checking in a book
	router.PATCH("/checkin", returnBook)


	//run the api ona  specific IP (in this case local host) and port (we have chose port 8080)
	router.Run("localhost:8080")

}
