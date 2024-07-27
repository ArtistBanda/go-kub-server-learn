package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", pingPong)

	router.Run(":8080")
}

func pingPong(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "pong")
}
