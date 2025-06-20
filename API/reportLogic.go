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

	for _, v := range activity {

		dur, err := parseTime(v.Duration)
		if err != nil {
			return err
		}
		totalTime += dur

		month, err := parseMonth(v.StartTime)
		if err != nil {
			return err
		}
		timeByMonth[month] += dur.Minutes()
	}

	report.TotalWatchTime = totalTime.Minutes()

	var months []Month
	for k, v := range timeByMonth {
		var m Month
		m.Month = k
		m.TimeSpent = int(math.Round(v))
		months = append(months, m)
	}

	report.Trends = months
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
