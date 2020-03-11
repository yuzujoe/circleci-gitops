package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, "Hello World")
	})

	log.Fatal(r.Run())
}