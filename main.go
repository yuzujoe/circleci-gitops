package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.JSON(200, "GitOps Like!!")
	})

	r.GET("/app-a", func(c *gin.Context) {
		c.JSON(200, "sample-app-a")
	})

	log.Fatal(r.Run(":8080"))
}
