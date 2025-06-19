package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tmdb "github.com/cyruzin/golang-tmdb"
)

func initializeTMDB() (*tmdb.Client, error) {
	apiKey := os.Getenv("TMDB_API_KEY")

	tmdbClient, err := tmdb.Init(apiKey)
	tmdbClient.SetClientAutoRetry()

	if err != nil {
		return nil, err
	}

	return tmdbClient, nil
}

func callTMDB(tmdbClient *tmdb.Client, title string) error {

	var counter int8
	var id int64
	// Check if the title contains a colon, indicating it might be a TV show
	if strings.Contains(title, ":") {

		seriesTitle := strings.Split(title, ": ")[0]
		log.Println("Searching for TV Show:", seriesTitle)

		searchResult, err := tmdbClient.GetSearchTVShow(seriesTitle, nil)
		if err != nil {
			return err
		}

		for _, result := range searchResult.Results {
			if result.Name == seriesTitle {
				id = result.ID
				counter++
				fmt.Println("Found TV Show:", result)
			}
		}
	} else
	// If the title does not contain a colon, treat it as a movie
	{
		log.Println("Searching for movie:", title)
		searchResult, err := tmdbClient.GetSearchMovies(title, nil)
		if err != nil {
			return err
		}
		for _, result := range searchResult.Results {
			if result.Title == title {
				id = result.ID
				counter++
			}
		}
	}

	if counter == 1 {
		log.Println(id)
	}

	// 	options := map[string]string{
	// 	"append_to_response": "images,credits",
	// }
	// tvDetails, err := tmdbClient.GetTVDetails(61222, nil) // ID BoJack Horseman
	// if err != nil {
	// 	log.Fatalf("Błąd: %v", err)
	// }

	// fmt.Println("Tytuł:", tvDetails.Name)
	// fmt.Println("Gatunki:")
	// for _, genre := range tvDetails.Genres {
	// 	fmt.Println("-", genre.Name)
	// }

	return nil
}
