package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
)

type book struct{
	ID  string   `json: "id"`   //upper case makes it an exported field, but when serialising in json convert the field name to lower case
	Title string  `json: "title"`
	Author string  `json: "author"`
	Quantity int	`json: "quantity"`
}

//mock database
var books = []book{
				{ID: "1",  Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
				{ID: "2",  Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
				{ID: "3",  Title: "War And Peace", Author: "Leo Tolstoy", Quantity: 6},
}

//set up router to handle main end points of api



func getBooks(c*gin.Context){
		c.IndentedJSON(http.StatusOK, books)   //nicely indented json, the status that we are sending is status okay, the data that we are sending is "books". So we return a json obejct that has all of the books in it
}

func createBook(c*gin.Context){ //c stores query parameters, headers

	var newBook book   	// the new book is of type book

	if err := c.BindJSON(&newBook); err != nil{

			 // if the error is not equal to null, in that case we shall simpy retur
				return
			}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}



func main(){
	router := gin.Default()  //create a gin router
	router.GET("/books", getBooks)   //when visting the /books endpoint the getbooks function is called

	//can route a specific route to a function with router variable


	//create a book:
	router.POST("/books", createBook)




	router.Run("localhost:8080")   //hte server and port we are running the web server on


}


//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
