package main

type ViewingActivity struct {
	StartTime      string `json:"startTime"`
	Duration       string `json:"duration"`
	Attributes     string `json:"attributes"`
	Title          string `json:"title"`
	DeviceType     string `json:"deviceType"`
	Bookmark       string `json:"bookmark"`
	LatestBookmark string `json:"latestBookmark"`
	Country        string `json:"country"`
}

type Profile struct {
	Name            string            `json:"name"`
	ViewingActivity []ViewingActivity `json:"viewingActivity"`
}

type Data struct {
	Profiles []Profile `json:"profiles"`
}
