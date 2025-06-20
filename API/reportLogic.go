package main

import (
	"fmt"
	"strings"
	"time"
)

func parseDuration(hms string) (time.Duration, error) {
	parts := strings.Split(hms, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("Incorrect time format: %s", hms)
	}

	return time.ParseDuration(fmt.Sprintf("%sh%sm%ss", parts[0], parts[1], parts[2]))
}

func analyseData(activity []ViewingActivity, report *Report) error {

	var total time.Duration

	for _, v := range activity {

		dur, err := parseDuration(v.Duration)
		if err != nil {
			return err
		}
		total += dur
	}

	report.TotalWatchTime = total.Minutes()

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
