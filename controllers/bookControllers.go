package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Books struct {
	BookId string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Books{}

// method untuk Get All Book
// dengan response
/*
[
	{
		"id" : 1,
		"title" : "Golang",
		"author" : "Gopher",
		"desc" : "A book for Go"
	}
]
*/
func GetAllBook(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, BookDatas)
}

// method untuk Get Book by Id
// dengan response
/*

	{
		"id" : 1,
		"title" : "Golang",
		"author" : "Gopher",
		"desc" : "A book for Go"
	}

*/
func GetBookByID(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false

	var bookData Books

	for i, book := range BookDatas {
		if bookID == book.BookId {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
			"error_status":  "Data Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, bookData)
}

// method Add Book
// response
// "Created"
func CreateBook(ctx *gin.Context) {
	var newBook Books

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookId = fmt.Sprintf("%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, "Created")

}

// method Update Book

func UpdatedBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var updatedBook Books

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	for i, book := range BookDatas {
		if bookID == book.BookId {
			condition = true
			BookDatas[i] = updatedBook
			BookDatas[i].BookId = bookID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, "Updated")
}

// method Delete Book

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if bookID == book.BookId {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Books{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, "Deleted")
}
