package main

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
