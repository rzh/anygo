package main

import "github.com/gin-gonic/gin"

func index(c *gin.Context) {
}

func posting(c *gin.Context) {

	// Parse JSON
	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if c.Bind(&json) == nil {
		c.JSON(200, gin.H{"status": "ok"})
	}
}

func main() {
	router := gin.Default()

	router.GET("/", index)
	router.POST("/somePost", posting)

	router.Run(":8080")
}
