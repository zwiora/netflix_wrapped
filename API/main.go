package main

import (
	"fmt"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/gin-gonic/gin"
)

var tmdbClient *tmdb.Client

func main() {
	var err error

	setApiKey()
	tmdbClient, err = initializeTMDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting application")
	port := "8080"

	router := gin.Default()
	router.GET("/test", getAllData)
	router.POST("/generate", postData)
	router.Run("localhost:" + port)

	fmt.Println("Server running on port " + port)
}
