package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type entity struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

var entities = []entity{
	{ID: "1", Content: "Test1"},
	{ID: "2", Content: "Test2"},
	{ID: "3", Content: "Test3"},
}

func getEntities(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, entities)
}

func main() {
	fmt.Println("Starting application")
	port := "8080"

	router := gin.Default()
	router.GET("/test", getEntities)
	router.Run("localhost:" + port)

	fmt.Println("Server running on port " + port)
}
