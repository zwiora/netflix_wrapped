package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// setApiKey()
	// tmdbClient, err := initializeTMDB()
	// if err != nil {
	// 	panic(err)
	// }
	// callTMDB(tmdbClient, "Friends: Hello: Episode 1")

	fmt.Println("Starting application")
	port := "8080"

	router := gin.Default()
	router.GET("/test", getAllData)
	router.POST("/generate", postData)
	router.Run("localhost:" + port)

	fmt.Println("Server running on port " + port)
}
