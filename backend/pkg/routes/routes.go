package routes

import (
	"github.com/Sattuu21/go-bookstore/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterBookstoreRoutes(r *gin.Engine) {
	r.POST("/book/", controllers.CreateBook)
	r.GET("/book/", controllers.GetBook)
	r.GET("/book/:bookId", controllers.GetBookById)
	r.DELETE("/book/:bookId", controllers.DeleteBook)
	r.PUT("/book/:bookId", controllers.UpdateBook)
}
