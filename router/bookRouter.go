package router

import (
	"assigment_5/restAPI/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)

	router.GET("/books", controllers.GetAllBook)

	router.GET("/books/:bookID", controllers.GetBookByID)

	router.PUT("/books/:bookID", controllers.UpdatedBook)

	router.DELETE("/books/:bookID", controllers.DeleteBook)

	return router
}
