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

type ProductionType string

const (
	Movie ProductionType = "Movie"
	TV    ProductionType = "TV series"
)

type Production struct {
	Title       string         `json:"title"`
	Type        ProductionType `json:"type"`
	Genre       string         `json:"genre"`
	Rating      float64        `json:"rating"`
	Duration    int            `json:"duration"`
	WatchedTime int            `json:"watchedTime"`
}

type Report struct {
	TopGenres      []string     `json:"topGenres"`
	TopActors      []string     `json:"topActors"`
	TotalWatchTime int          `json:"totalWatchTime"`
	BingeData      []int        `json:"bingeData"`
	GenresData     []string     `json:"genresData"`
	TrendsData     []string     `json:"trendsData"`
	Watched        []Production `json:"watched"`
}

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

func main() {

	fmt.Println("Starting application")
	port := "8080"

	router := gin.Default()
	router.GET("/test", getAllData)
	router.POST("/generate", postData)
	router.Run("localhost:" + port)

	fmt.Println("Server running on port " + port)
}
