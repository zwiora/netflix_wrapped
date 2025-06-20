package main

type ProductionType string

const (
	Movie ProductionType = "Movie"
	TV    ProductionType = "TV series"
)

type Production struct {
	Title       string         `json:"title"`
	Type        ProductionType `json:"type"`
	Genre       []string       `json:"genre"`
	Rating      float64        `json:"rating"`
	WatchedTime int            `json:"watchedTime"`
}

type ProductionDetailed struct {
	Title            string         `json:"title"`
	Type             ProductionType `json:"type"`
	Genre            []string       `json:"genre"`
	Rating           float64        `json:"rating"`
	WatchedTime      int            `json:"watchedTime"`
	ImageURL         string         `json:"imageUrl"`
	Overview         string         `json:"overview"`
	FullDuration     int            `json:"fullRuntime"`
	FirstAirDate     string         `json:"firstAirDate"`
	NumberOfSeasons  int            `json:"numberOfSeasons"`
	NumberOfEpisodes int            `json:"numberOfEpisodes"`
	Adult            bool           `json:"adult"`
	OriginCountry    []string       `json:"originCountry"`
	OriginalLanguage string         `json:"originalLanguage"`
	OriginalTitle    string         `json:"originalTitle"`
	Recommended      []Production   `json:"recommended"`
}

type Actor struct {
	Name            string       `json:"name"`
	AlsoKnownAs     []string     `json:"alsoKnownAs"`
	Gender          string       `json:"gender"`
	KnownFor        []Production `json:"knownFor"`
	DateOfBirth     string       `json:"dateOfBirth"`
	DateOfDeath     string       `json:"dateOfDeath"`
	PlaceOfBirth    string       `json:"placeOfBirth"`
	Biography       string       `json:"biography"`
	ProfileImageURL string       `json:"profileImageUrl"`
	YourProductions []Production `json:"yourProductions"`
}

type Genre struct {
	Name                string `json:"name"`
	NumberOfProductions int    `json:"numberOfProductions"`
	TimeSpent           int    `json:"timeSpent"` // in minutes
}

type BingeSession struct {
	Production        Production `json:"production"`
	StartTime         string     `json:"startTime"`
	EndTime           string     `json:"endTime"`
	TotalWatchTime    int        `json:"totalWatchTime"` // in minutes
	NumberOfEpisodes  int        `json:"numberOfEpisodes"`
	PercentageWatched float64    `json:"percentageWatched"`
}

type Month struct {
	Number    int `json:"number"`    // 1-12 for January-December
	TimeSpent int `json:"timeSpent"` // in minutes
}

type Report struct {
	UserName        string             `json:"userName"`
	TotalWatchTime  float64            `json:"totalWatchTime"`
	AverageRating   float64            `json:"averageRating"`
	BestMovie       ProductionDetailed `json:"bestMovie"`
	BestTV          ProductionDetailed `json:"bestTV"`
	WorstMovie      ProductionDetailed `json:"worstMovie"`
	WorstTV         ProductionDetailed `json:"worstTV"`
	FavouriteActors []Actor            `json:"favouriteActors"`
	FavouriteGenres []Genre            `json:"favouriteGenres"`
	BingeSessions   []BingeSession
	Trends          []Month      `json:"trends"`
	WatchedList     []Production `json:"watchedList"`
}
