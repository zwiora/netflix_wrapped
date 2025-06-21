package main

import (
	"fmt"
	"log"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/gin-gonic/gin"
)

var tmdbClient *tmdb.Client

func main() {
	// External API
	var err error
	setApiKey()
	tmdbClient, err = initializeTMDB()
	if err != nil {
		log.Println(err)
	}

	// API service
	fmt.Println("Starting application")
	port := "8080"

	router := gin.Default()
	router.POST("/generate", postData)
	router.Run("localhost:" + port)
}
