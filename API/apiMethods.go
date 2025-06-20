package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var mockReport = Report{
// 	TotalWatchTime: 123,
// 	AverageRating:  7.9,
// 	BestMovie: ProductionDetailed{
// 		Title:            "Inception",
// 		Type:             Movie,
// 		Genre:            []string{"Sci-Fi", "Thriller"},
// 		Rating:           9.0,
// 		WatchedTime:      148,
// 		ImageURL:         "https://image.tmdb.org/t/p/w500/inception.jpg",
// 		Overview:         "A thief who steals corporate secrets through the use of dream-sharing technology...",
// 		FullDuration:     148,
// 		FirstAirDate:     "2010-07-16",
// 		NumberOfSeasons:  0,
// 		NumberOfEpisodes: 0,
// 		Adult:            false,
// 		OriginCountry:    []string{"US"},
// 		OriginalLanguage: "en",
// 		OriginalTitle:    "Inception",
// 		Recommended: []Production{
// 			{Title: "Interstellar", Type: Movie, Genre: []string{"Sci-Fi"}, Rating: 8.6, WatchedTime: 169},
// 			{Title: "Tenet", Type: Movie, Genre: []string{"Action", "Sci-Fi"}, Rating: 7.4, WatchedTime: 150},
// 		},
// 	},
// 	BestTV: ProductionDetailed{
// 		Title:            "Breaking Bad",
// 		Type:             TV,
// 		Genre:            []string{"Crime", "Drama"},
// 		Rating:           9.5,
// 		WatchedTime:      300,
// 		ImageURL:         "https://image.tmdb.org/t/p/w500/breakingbad.jpg",
// 		Overview:         "A high school chemistry teacher turned methamphetamine producer...",
// 		FullDuration:     3000,
// 		FirstAirDate:     "2008-01-20",
// 		NumberOfSeasons:  5,
// 		NumberOfEpisodes: 62,
// 		Adult:            false,
// 		OriginCountry:    []string{"US"},
// 		OriginalLanguage: "en",
// 		OriginalTitle:    "Breaking Bad",
// 		Recommended: []Production{
// 			{Title: "Better Call Saul", Type: TV, Genre: []string{"Crime", "Drama"}, Rating: 8.7, WatchedTime: 150},
// 			{Title: "Ozark", Type: TV, Genre: []string{"Crime"}, Rating: 8.4, WatchedTime: 120},
// 		},
// 	},
// 	WorstMovie: ProductionDetailed{
// 		Title:            "The Room",
// 		Type:             Movie,
// 		Genre:            []string{"Drama"},
// 		Rating:           3.2,
// 		WatchedTime:      99,
// 		ImageURL:         "https://image.tmdb.org/t/p/w500/theroom.jpg",
// 		Overview:         "A melodramatic love triangle full of awkward dialogue...",
// 		FullDuration:     99,
// 		FirstAirDate:     "2003-06-27",
// 		NumberOfSeasons:  0,
// 		NumberOfEpisodes: 0,
// 		Adult:            true,
// 		OriginCountry:    []string{"US"},
// 		OriginalLanguage: "en",
// 		OriginalTitle:    "The Room",
// 		Recommended:      []Production{},
// 	},
// 	WorstTV: ProductionDetailed{
// 		Title:            "Velma",
// 		Type:             TV,
// 		Genre:            []string{"Animation", "Comedy"},
// 		Rating:           1.9,
// 		WatchedTime:      45,
// 		ImageURL:         "https://image.tmdb.org/t/p/w500/velma.jpg",
// 		Overview:         "The origin story of Velma Dinkley...",
// 		FullDuration:     120,
// 		FirstAirDate:     "2023-01-12",
// 		NumberOfSeasons:  1,
// 		NumberOfEpisodes: 10,
// 		Adult:            false,
// 		OriginCountry:    []string{"US"},
// 		OriginalLanguage: "en",
// 		OriginalTitle:    "Velma",
// 		Recommended:      []Production{},
// 	},
// 	FavouriteActors: []Actor{
// 		{
// 			Name:            "Bryan Cranston",
// 			AlsoKnownAs:     []string{"Walter White"},
// 			Gender:          "Male",
// 			KnownFor:        []Production{{Title: "Breaking Bad", Type: TV, Genre: []string{"Drama"}, Rating: 9.5, WatchedTime: 300}},
// 			DateOfBirth:     "1956-03-07",
// 			DateOfDeath:     "",
// 			PlaceOfBirth:    "Hollywood, California, USA",
// 			Biography:       "Bryan Cranston is an American actor, best known for his role as Walter White...",
// 			ProfileImageURL: "https://image.tmdb.org/t/p/w500/bryan.jpg",
// 			YourProductions: []Production{{Title: "Breaking Bad", Type: TV, Genre: []string{"Drama"}, Rating: 9.5, WatchedTime: 300}},
// 		},
// 		{
// 			Name:            "Leonardo DiCaprio",
// 			AlsoKnownAs:     []string{"Leo"},
// 			Gender:          "Male",
// 			KnownFor:        []Production{{Title: "Inception", Type: Movie, Genre: []string{"Sci-Fi"}, Rating: 9.0, WatchedTime: 148}},
// 			DateOfBirth:     "1974-11-11",
// 			DateOfDeath:     "",
// 			PlaceOfBirth:    "Los Angeles, California, USA",
// 			Biography:       "Oscar-winning actor known for Titanic, Inception, and The Revenant...",
// 			ProfileImageURL: "https://image.tmdb.org/t/p/w500/leo.jpg",
// 			YourProductions: []Production{{Title: "Inception", Type: Movie, Genre: []string{"Sci-Fi"}, Rating: 9.0, WatchedTime: 148}},
// 		},
// 	},
// 	FavouriteGenres: []Genre{
// 		{Name: "Drama", NumberOfProductions: 12, TimeSpent: 1040},
// 		{Name: "Sci-Fi", NumberOfProductions: 5, TimeSpent: 620},
// 		{Name: "Crime", NumberOfProductions: 7, TimeSpent: 920},
// 	},
// 	BingeSessions: []BingeSession{
// 		{
// 			Production:        Production{Title: "Breaking Bad", Type: TV, Genre: []string{"Crime"}, Rating: 9.5, WatchedTime: 200},
// 			StartTime:         "2025-04-01T17:00:00",
// 			EndTime:           "2025-04-01T22:00:00",
// 			TotalWatchTime:    300,
// 			NumberOfEpisodes:  6,
// 			PercentageWatched: 85.0,
// 		},
// 		{
// 			Production:        Production{Title: "Stranger Things", Type: TV, Genre: []string{"Sci-Fi"}, Rating: 8.6, WatchedTime: 180},
// 			StartTime:         "2025-05-03T19:00:00",
// 			EndTime:           "2025-05-03T23:00:00",
// 			TotalWatchTime:    240,
// 			NumberOfEpisodes:  4,
// 			PercentageWatched: 75.0,
// 		},
// 	},
// 	Trends: []Month{
// 		{Month: "2024.11", TimeSpent: 320},
// 		{Month: "2024.12", TimeSpent: 410},
// 		{Month: "2025.01", TimeSpent: 550},
// 		{Month: "2025.02", TimeSpent: 610},
// 		{Month: "2025.03", TimeSpent: 720},
// 	},
// 	WatchedList: []Production{
// 		{Title: "Inception", Type: Movie, Genre: []string{"Sci-Fi"}, Rating: 9.0, WatchedTime: 148},
// 		{Title: "Breaking Bad", Type: TV, Genre: []string{"Crime", "Drama"}, Rating: 9.5, WatchedTime: 300},
// 		{Title: "The Room", Type: Movie, Genre: []string{"Drama"}, Rating: 3.2, WatchedTime: 99},
// 		{Title: "Velma", Type: TV, Genre: []string{"Animation"}, Rating: 1.9, WatchedTime: 45},
// 	},
// }

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

	report, err := generateReport(&newEntity)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "An error ocurred",
		})
	} else {
		uploadedData = append(uploadedData, newEntity)
		c.IndentedJSON(http.StatusOK, report)
	}
}
