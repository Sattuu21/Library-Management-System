package main

import (
	"fmt"
	"log"

	"github.com/Sattuu21/go-bookstore/pkg/routes"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Apply CORS middleware
	r.Use(CORSMiddleware())

	// Register bookstore routes
	routes.RegisterBookstoreRoutes(r)

	fmt.Println("Listening and serving on port 9010")
	if err := r.Run("localhost:9010"); err != nil {
		log.Fatal("Server Run Failed: ", err)
	}
}
