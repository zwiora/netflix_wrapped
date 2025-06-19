package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var uploadedData = []Data{}

func getAllData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, uploadedData)
}

func postData(c *gin.Context) {
	var newEntity Data
	example := Report{
		TopGenres:      []string{"Drama", "Comedy", "Sci-Fi"},
		TopActors:      []string{"Bryan Cranston", "Emma Stone"},
		TotalWatchTime: 920,
		BingeData:      []int{3, 5, 2, 4},
		GenresData:     []string{"Drama", "Thriller", "Comedy"},
		TrendsData:     []string{"Popular", "Trending", "Critically Acclaimed"},
		Watched: []Production{
			{
				Title:       "Breaking Bad",
				Type:        TV,
				Genre:       "Drama",
				Rating:      9.5,
				Duration:    47,
				WatchedTime: 47,
			},
			{
				Title:       "La La Land",
				Type:        Movie,
				Genre:       "Musical",
				Rating:      8.0,
				Duration:    128,
				WatchedTime: 128,
			},
			{
				Title:       "Stranger Things",
				Type:        TV,
				Genre:       "Sci-Fi",
				Rating:      8.7,
				Duration:    50,
				WatchedTime: 50,
			},
		},
	}

	if err := c.BindJSON(&newEntity); err != nil {
		log.Println(err)
		return
	}

	uploadedData = append(uploadedData, newEntity)
	c.IndentedJSON(http.StatusOK, example)
}
