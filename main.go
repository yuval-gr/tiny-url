package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", getTinyURLs)
	router.GET("/:tiny", getURL)
	router.POST("/", createTinyURL)

	router.Run("0.0.0.0:8080")
}
