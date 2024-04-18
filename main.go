package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/home", testRoute)

	router.Run(":8080")
}

func welcome(s string) string {
	return fmt.Sprintf("Welcome, %s", s)
}

func testRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}
