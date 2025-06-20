package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

func parseTime(timeStr string) (time.Duration, error) {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("Incorrect time format: %s", timeStr)
	}

	return time.ParseDuration(fmt.Sprintf("%sh%sm%ss", parts[0], parts[1], parts[2]))
}

func parseMonth(dateStr string) (string, error) {
	layout := "1/2/2006 15:04"

	t, err := time.Parse(layout, dateStr)

	if err != nil {
		layout := "2006-01-02 15:04:05"
		t, err = time.Parse(layout, dateStr)
		if err != nil {

			return "", err
		}
	}

	return strconv.Itoa(t.Year()) + "." + strconv.Itoa(int(t.Month())), nil
}

func IsWithinThreeHours(dateStr1, dateStr2 string) (bool, error) {
	layouts := []string{
		"1/2/2006 15:04",      // m/d/y h:mm
		"2006-01-02 15:04:05", // ISO format
	}

	var t1, t2 time.Time
	var err error

	// first date
	for _, layout := range layouts {
		t1, err = time.Parse(layout, dateStr1)
		if err == nil {
			break
		}
	}
	if err != nil {
		return false, err
	}

	// second date
	err = nil
	for _, layout := range layouts {
		t2, err = time.Parse(layout, dateStr2)
		if err == nil {
			break
		}
	}
	if err != nil {
		return false, err
	}

	diff := t1.Sub(t2)
	if diff < 0 {
		diff = -diff
	}

	return diff <= 3*time.Hour, nil
}

func generateProdDetails(production *Production) *ProductionDetailed {
	result := new(ProductionDetailed)

	result.Genre = production.Genre
	result.Title = production.Title
	result.Rating = production.Rating
	result.WatchedTime = production.WatchedTime
	result.Type = production.Type
	result.id = production.id
	getProductionDetailsTMDB(result)

	return result
}

func generateList(list map[string]*Production, genres map[int64]string) ([]Production, error) {
	var result []Production

	for _, v := range list {
		id, rating, genresIDs, err := getProductionTMDB(v.Title, v.Type)
		if err != nil {
			return nil, err
		}

		v.id = int(id)
		v.Rating = rating
		v.WatchedTime = float32(math.Round(float64(v.WatchedTime)))

		for _, r := range genresIDs {
			v.Genre = append(v.Genre, genres[r])
		}

		result = append(result, *v)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].WatchedTime > result[j].WatchedTime
	})

	return result, nil
}

func analyseData(activity []ViewingActivity, report *Report) error {

	var err error
	var totalTime time.Duration
	timeByMonth := make(map[string]float64)
	watchedMovies := make(map[string]*Production)
	watchedTV := make(map[string]*Production)
	bingeSessions := make(map[string][]BingeSession)

	for _, v := range activity {
		dur, err := parseTime(v.Duration)
		if err != nil {
			return err
		}

		if dur.Minutes() >= 5 {
			// total time
			totalTime += dur

			// trends
			month, err := parseMonth(v.StartTime)
			if err != nil {
				return err
			}
			timeByMonth[month] += dur.Minutes()

			// production
			var production *Production
			var exists bool
			title := v.Title
			splited := strings.Split(title, ": ")

			if len(splited) > 2 {
				title = splited[0]

				production, exists = watchedTV[title]
				if !exists {
					watchedTV[title] = new(Production)
					production = watchedTV[title]

					production.Title = title
				}

				production.Type = TV

				// binge session

				session, exists := bingeSessions[title]
				if !exists {
					var s BingeSession
					s.Production = production
					s.episodes = []string{v.Title}
					s.StartTime = v.StartTime
					s.EndTime = v.StartTime
					s.TotalWatchTime = float32(dur.Minutes())

					bingeSessions[title] = []BingeSession{s}
				} else {
					s := session[len(session)-1]
					isWithin, err := IsWithinThreeHours(s.StartTime, v.StartTime)
					if err != nil {
						return err
					}
					if isWithin {
						if !slices.Contains(s.episodes, v.Title) {
							s.episodes = append(s.episodes, v.Title)
						}
						s.StartTime = v.StartTime
						s.TotalWatchTime += float32(dur.Minutes())

						session[len(session)-1] = s
						bingeSessions[title] = session
					} else {
						var b BingeSession
						b.Production = production
						b.episodes = []string{v.Title}
						b.StartTime = v.StartTime
						b.EndTime = v.StartTime
						b.TotalWatchTime = float32(dur.Minutes())

						bingeSessions[title] = append(bingeSessions[title], b)
					}
				}
			} else {
				production, exists = watchedMovies[title]
				if !exists {
					watchedMovies[title] = new(Production)
					production = watchedMovies[title]

					production.Title = title
				}
				production.Type = Movie
			}
			production.WatchedTime += float32(dur.Minutes())
		}
	}

	// total time
	report.TotalWatchTime = totalTime.Minutes()

	// trends
	var months []Month
	for k, v := range timeByMonth {
		var m Month
		m.Month = k
		m.TimeSpent = int(math.Round(v))
		months = append(months, m)
	}
	report.Trends = months

	// genres
	genres, err := getGenresTMDB()
	if err != nil {
		return err
	}

	// productions
	log.Println("Generating list of movies")
	report.WatchedMovies, err = generateList(watchedMovies, genres)
	log.Println("Generating list of TV series")
	report.WatchedTV, err = generateList(watchedTV, genres)

	if err != nil {
		return err
	}

	// binge
	var resultBinges []BingeSession
	for _, list := range bingeSessions {
		for _, session := range list {
			if len(session.episodes) >= 3 {
				session.NumberOfWatchedEpisodes = len(session.episodes)
				session.TotalWatchTime = float32(math.Round(float64(session.TotalWatchTime)))
				resultBinges = append(resultBinges, session)
			}
		}
	}
	sort.Slice(resultBinges, func(i, j int) bool {
		return resultBinges[i].TotalWatchTime > resultBinges[j].TotalWatchTime
	})
	report.BingeSessions = resultBinges

	// Second traversal
	var counter int
	var sum float32
	timeByGenre := make(map[string]float32)
	productionsByGenre := make(map[string]int)
	var bestMovie Production
	var bestTV Production
	var worstMovie Production
	var worstTV Production
	bestMovie.Rating = 0
	bestTV.Rating = 0
	worstMovie.Rating = 11
	worstTV.Rating = 11

	for _, v := range report.WatchedMovies {
		//average rating
		counter++
		sum += v.Rating

		//genre
		for _, g := range v.Genre {
			timeByGenre[g] += v.WatchedTime
			productionsByGenre[g]++
		}

		//ranking
		if bestMovie.Rating < v.Rating {
			bestMovie = v
		}
		if worstMovie.Rating > v.Rating {
			worstMovie = v
		}
	}
	for _, v := range report.WatchedTV {
		//average rating
		counter++
		sum += v.Rating

		//genre
		for _, g := range v.Genre {
			timeByGenre[g] += v.WatchedTime
			productionsByGenre[g]++
		}

		//ranking
		if bestTV.Rating < v.Rating {
			bestTV = v
		}
		if worstTV.Rating > v.Rating {
			worstTV = v
		}
	}

	// average rating
	report.AverageRating = sum / float32(counter)

	// genre
	var genresArr []Genre
	for k, v := range timeByGenre {
		g := new(Genre)
		g.Name = k
		g.TimeSpent = int(math.Round(float64(v)))
		g.NumberOfProductions = productionsByGenre[k]

		genresArr = append(genresArr, *g)
	}
	sort.Slice(genresArr, func(i, j int) bool {
		return genresArr[i].TimeSpent > genresArr[j].TimeSpent
	})
	report.Genres = genresArr

	// ranking
	report.BestMovie = *generateProdDetails(&bestMovie)
	report.BestTV = *generateProdDetails(&bestTV)
	report.WorstMovie = *generateProdDetails(&worstMovie)
	report.WorstTV = *generateProdDetails(&worstTV)

	return nil
}

func generateReport(data *Data) (*Report, error) {
	var report *Report
	report = new(Report)

	profileIdx := 2
	userData := data.Profiles[profileIdx]

	// fmt.Println(userData.N)

	report.UserName = userData.Name
	err := analyseData(userData.ViewingActivity, report)

	return report, err
}
