package main
import (
"net/http"
"github.com/gin-gonic/gin"
"errors"
"strconv"
)

type book struct{
	ID int	`json:"id"`
	Title string	`json:"title"`
	Author string	`json:"author"`
	Quantity int	`json:"quantity"`
}


var books = []book{
	{ID:1,Title:"In search of lost time", Author:"Marcel Proust", Quantity:2},
	{ID:2,Title:"Beyond good and evil", Author:"Fredrick niechze", Quantity:5},
	{ID:3,Title:"The brothers Karamazov", Author:"Fyodor Dostoevesky", Quantity:1},
	{ID:4,Title:"Metamorphosis", Author:"Frank Kafka", Quantity:9},
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK,books)

}

func CreateBook(c *gin.Context){
	var newBook book 
	if err := c.BindJSON(&newBook); err !=nil{
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated,newBook)
}

func bookbyID(c *gin.Context){
	strid := c.Param("id")
	id,err := strconv.Atoi(strid)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	book,err := GetBookid(id)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK,book)
}

func GetBookid(id int)(*book, error){
	for i, b := range books{
		if b.ID == id{
			return &books[i],nil
		}
	}
	return nil, errors.New("book not found")
}

func checkoutbook(c *gin.Context){
	strid,ok := c.GetQuery("id")
	if !ok{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"missing query parameter"})
		return
	}
	bookid, err := strconv.Atoi(strid)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid request"})
		return
	}
	book,err := GetBookid(bookid)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
        return
	}
	if book.Quantity <= 0{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"missing query parameter"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK,book)
}

func returnBook(c *gin.Context){
	strid,ok := c.GetQuery("id")
	if !ok{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"missing query parameter"})
		return
	}
	bookid, err := strconv.Atoi(strid)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid request"})
		return
	}
	book,err := GetBookid(bookid)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
        return
	}
	book.Quantity +=1
	c.IndentedJSON(http.StatusOK,book)
}

func main (){
	router := gin.Default()
	router.GET("/books",getBooks)
	router.GET("/books/:id",bookbyID)
	router.POST("/books",CreateBook)
	router.PATCH("/cheackout",checkoutbook)
	router.PATCH("/return",returnBook)
	router.Run("localhost:8080")
}