package main

import (
    "github.com/dev-kvt/restfulgo/api"
	"github.com/gin-gonic/gin"
)

func main(){
	api.Initialze()
	r:=gin.Default()
	r.POST("/token", api.)
	// protected routes
	protected := r.Group("/", api.JWTAuthMiddleware())

	//routes
	r.POST("/book", api.CreateBook)
	r.GET("/books", api.GetBooks)
	r.GET("/book/:id", api.GetBook)
	r.PUT("/book/:id", api.UpdateBook)
	r.DELETE("/book/:id", api.DeleteBook)
	r.Run("localhost:8080")
}