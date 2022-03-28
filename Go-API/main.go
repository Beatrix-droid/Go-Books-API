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

//mock database
var books = []book{{
										ID: "1",  Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
										{ID: "2",  Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
										{ID: "3",  Title: "War And Peace", Author: "Leo Tolstoy", Quantity: 6},
}

//set up router to handle main end points of api



func getBooks(c*gin.Context){
		c.IndentedJson(http.StatusOK, books)   //nicely indented json, the status that we are sending is status okay, the data that we are sending is "books". So we return a json obejct that has all of the books in it
		router.GET("/books", getBooks)   //when visting the /books endpoint the getbooks function is called
		router.Run("localhost:8080")   //hte server and port we are running the web server on

}



func main(){
	router := gin.Default()  //create a gin router
	//can route a specific route to a function with router variable



}
