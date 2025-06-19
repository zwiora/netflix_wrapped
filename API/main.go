package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewingActivity struct {
	StartTime      string `json:"startTime"`
	Duration       string `json:"duration"`
	Attributes     string `json:"attributes"`
	Title          string `json:"title"`
	DeviceType     string `json:"deviceType"`
	Bookmark       string `json:"bookmark"`
	LatestBookmark string `json:"latestBookmark"`
	Country        string `json:"country"`
}

type Profile struct {
	Name            string            `json:"name"`
	ViewingActivity []ViewingActivity `json:"viewingActivity"`
}

type Data struct {
	Profiles []Profile `json:"profiles"`
}

var uploadedData = []Data{}

func getAllData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, uploadedData)
}

func postData(c *gin.Context) {
	var newEntity Data

	if err := c.BindJSON(&newEntity); err != nil {
		log.Println(err)
		return
	}

	uploadedData = append(uploadedData, newEntity)
	c.IndentedJSON(http.StatusCreated, newEntity)
}

func main() {
	fmt.Println("Starting application")
	port := "8080"

	router := gin.Default()
	router.GET("/test", getAllData)
	router.POST("/generate", postData)
	router.Run("localhost:" + port)

	fmt.Println("Server running on port " + port)
}
