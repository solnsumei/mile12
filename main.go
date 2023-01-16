package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/config"
)

func init() {
	config.LoadConfig()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Mile12 API",
		})
	})

	router.Run()
}
