package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello Argo Rollout")
	})

	r.GET("/app-a", func(c *gin.Context) {
		c.JSON(200, "Hello World")
	})

	log.Fatal(r.Run())
}
