package main

import (
	"fmt"
	"math"
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

func analyseData(activity []ViewingActivity, report *Report) error {

	var totalTime time.Duration
	timeByMonth := make(map[string]float64)
	watchedMovies := make(map[string]*Production)
	watchedTV := make(map[string]*Production)

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
			} else {
				production, exists = watchedMovies[title]
				if !exists {
					watchedMovies[title] = new(Production)
					production = watchedMovies[title]

					production.Title = title
				}
				production.Type = Movie
			}
			production.WatchedTime += dur.Minutes()
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

	// productions
	var prodsTV []Production
	for _, v := range watchedTV {
		v.WatchedTime = math.Round(v.WatchedTime)
		prodsTV = append(prodsTV, *v)
	}
	report.WatchedTV = prodsTV

	var prodsM []Production
	for _, v := range watchedMovies {
		v.WatchedTime = math.Round(v.WatchedTime)
		prodsM = append(prodsM, *v)
	}
	report.WatchedMovies = prodsM

	return nil
}

func generateReport(data *Data) (*Report, error) {
	var report *Report
	report = new(Report)

	profileIdx := 0
	userData := data.Profiles[profileIdx]

	// fmt.Println(userData.N)

	report.UserName = userData.Name
	err := analyseData(userData.ViewingActivity, report)

	return report, err
}
