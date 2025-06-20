package main

import (
	"log"
	"os"

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

func callTMDB(title string, prodType ProductionType) (int64, float32, []int64, error) {

	var id int64
	var rating float32
	var genres []int64

	if prodType == TV {

		log.Println("Searching for TV Show:", title)

		searchResult, err := tmdbClient.GetSearchTVShow(title, nil)
		if err != nil {
			return 0, 0, nil, err
		}

		for _, result := range searchResult.Results {
			if result.Name == title {
				id = result.ID
				rating = result.VoteAverage
				genres = result.GenreIDs
				break
			}
		}
	} else {
		log.Println("Searching for movie:", title)
		searchResult, err := tmdbClient.GetSearchMovies(title, nil)
		if err != nil {
			return 0, 0, nil, err
		}
		for _, result := range searchResult.Results {
			if result.Title == title {
				id = result.ID
				rating = result.VoteAverage
				genres = result.GenreIDs
				break
			}
		}
	}

	return id, rating, genres, nil
	// options := map[string]string{
	// 	"append_to_response": "images,credits",
	// }
	// productionDetails, err := tmdbClient.GetTVDetails(int(id), options) // ID BoJack Horseman
	// if err != nil {
	// 	return err
	// }

	// // fmt.Println(productionDetails.TVImagesAppend.Images.Posters[0].FilePath)
	// fmt.Println(productionDetails)

	// fmt.Println("Tytuł:", tvDetails.Name)
	// fmt.Println("Gatunki:")

	// for _, genre := range tvDetails.Genres {
	// 	fmt.Println("-", genre.Name)
	// }

	// return nil
}
