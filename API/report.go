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
	Rating      float32        `json:"rating"`
	WatchedTime float32        `json:"watchedTime"`
	id          int
}

type ProductionDetailed struct {
	Title            string         `json:"title"`
	Type             ProductionType `json:"type"`
	Genre            []string       `json:"genre"`
	Rating           float32        `json:"rating"`
	WatchedTime      float32        `json:"watchedTime"`
	id               int
	ImageURL         string   `json:"imageUrl"`
	Overview         string   `json:"overview"`
	FullDuration     int      `json:"fullDuration"`
	FirstAirDate     string   `json:"firstAirDate"`
	NumberOfSeasons  int      `json:"numberOfSeasons"`
	NumberOfEpisodes int      `json:"numberOfEpisodes"`
	OriginCountry    []string `json:"originCountry"`
	OriginalLanguage string   `json:"originalLanguage"`
	OriginalTitle    string   `json:"originalTitle"`
}

type Genre struct {
	Name                string `json:"name"`
	NumberOfProductions int    `json:"numberOfProductions"`
	TimeSpent           int    `json:"timeSpent"`
}

type BingeSession struct {
	Production              *Production `json:"production"`
	StartTime               string      `json:"startTime"`
	EndTime                 string      `json:"endTime"`
	TotalWatchTime          float32     `json:"totalWatchTime"`
	NumberOfWatchedEpisodes int         `json:"numberOfWatchedEpisodes"`
	episodes                []string
}

type Month struct {
	Month     string `json:"month"`
	TimeSpent int    `json:"timeSpent"`
}

type Report struct {
	UserName       string             `json:"userName"`
	TotalWatchTime float64            `json:"totalWatchTime"`
	AverageRating  float32            `json:"averageRating"`
	BestMovie      ProductionDetailed `json:"bestMovie"`
	BestTV         ProductionDetailed `json:"bestTV"`
	WorstMovie     ProductionDetailed `json:"worstMovie"`
	WorstTV        ProductionDetailed `json:"worstTV"`
	Genres         []Genre            `json:"genres"`
	BingeSessions  []BingeSession     `json:"bingeSessions"`
	Trends         []Month            `json:"trends"`
	WatchedMovies  []Production       `json:"watchedMovies"`
	WatchedTV      []Production       `json:"watchedTV"`
}
