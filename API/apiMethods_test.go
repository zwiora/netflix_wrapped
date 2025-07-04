package main

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var mockGenerateReport func(data *Data) (Report, error)

func init() {
	var err error
	setApiKey()
	tmdbClient, err = initializeTMDB()
	if err != nil {
		log.Println(err)
	}

	mockGenerateReport = func(data *Data) (Report, error) {
		return mockGenerateReport(data)
	}
}

func setupRouter() *gin.Engine {

	r := gin.Default()
	r.POST("/generate", postData)
	return r
}

func TestPostData_BadJSON(t *testing.T) {
	router := setupRouter()
	body := `{"invalid": "json`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/generate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "ERROR: Sent data is incorrect (check if your JSON file is in the correct format)")
}

func TestPostData_EmptyData(t *testing.T) {
	router := setupRouter()
	body := `{"profiles": []}`

	mockGenerateReport = func(data *Data) (Report, error) {
		return Report{}, errors.New("mock error")
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/generate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "ERROR: Empty data")
}

func TestPostData_EmptyHistory(t *testing.T) {
	router := setupRouter()
	body := `{"profiles": [{"name": "TestUser", "viewingActivity": []}]}`

	mockGenerateReport = func(data *Data) (Report, error) {
		return Report{UserName: "TestUser"}, nil
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/generate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "ERROR: Empty history")
}

func TestPostData_NotEnoughData(t *testing.T) {
	router := setupRouter()
	body := `{"profiles": [{"name": "User1","viewingActivity": [{
                    "startTime": "7/28/2022 4:22",
                    "duration": "0:00:25",
                    "attributes": "Autoplayed: user action: None; ",
                    "title": "Season 1 Trailer: Heartstopper",
                    "deviceType": "Worlds Best TV",
                    "bookmark": "0:00:25",
                    "latestBookmark": "0:00:25",
                    "country": "US (United States)"
                }]}]}`

	mockGenerateReport = func(data *Data) (Report, error) {
		return Report{UserName: "TestUser"}, nil
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/generate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "ERROR: Not enough data")
}

func TestPostData_Success(t *testing.T) {
	router := setupRouter()
	body := `{"profiles": [{"name": "User1","viewingActivity": [{
                    "startTime": "1/3/2020 3:12",
                    "duration": "1:30:16",
                    "attributes": "Autoplayed: user action: Unspecified; ",
                    "title": "I'll See You in My Dreams",
                    "deviceType": "Worlds 4th Best TV",
                    "bookmark": "1:30:16",
                    "latestBookmark": "1:30:16",
                    "country": "US (United States)"
                }]}]}`

	mockGenerateReport = func(data *Data) (Report, error) {
		return Report{UserName: "User1"}, nil
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/generate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	expectedResult := "{\n    \"userName\": \"User1\",\n    \"totalWatchTime\": 90.26666666666667,\n    \"averageRating\": 6.293,\n    \"bestMovie\": {\n        \"title\": \"I'll See You in My Dreams\",\n        \"type\": \"Movie\",\n        \"genre\": [\n            \"Drama\",\n            \"Comedy\"\n        ],\n        \"rating\": 6.293,\n        \"watchedTime\": 90,\n        \"imageUrl\": \"/dBN2Z525tklv9DSS3v7Tf7FWLoI.jpg\",\n        \"overview\": \"A sudden loss disrupts Carol’s orderly life, propelling her into the dating world for the first time in 20 years. Finally living in the present tense, she finds herself swept up in not one, but two unexpected relationships that challenge her assumptions about what it means to grow old.\",\n        \"fullDuration\": 92,\n        \"firstAirDate\": \"2015-05-15\",\n        \"numberOfSeasons\": 0,\n        \"numberOfEpisodes\": 0,\n        \"originCountry\": [\n            \"US\"\n        ],\n        \"originalLanguage\": \"en\",\n        \"originalTitle\": \"I'll See You in My Dreams\"\n    },\n    \"bestTV\": {\n        \"title\": \"\",\n        \"type\": \"\",\n        \"genre\": null,\n        \"rating\": 0,\n        \"watchedTime\": 0,\n        \"imageUrl\": \"\",\n        \"overview\": \"\",\n        \"fullDuration\": 0,\n        \"firstAirDate\": \"\",\n        \"numberOfSeasons\": 0,\n        \"numberOfEpisodes\": 0,\n        \"originCountry\": null,\n        \"originalLanguage\": \"\",\n        \"originalTitle\": \"\"\n    },\n    \"worstMovie\": {\n        \"title\": \"I'll See You in My Dreams\",\n        \"type\": \"Movie\",\n        \"genre\": [\n            \"Drama\",\n            \"Comedy\"\n        ],\n        \"rating\": 6.293,\n        \"watchedTime\": 90,\n        \"imageUrl\": \"/dBN2Z525tklv9DSS3v7Tf7FWLoI.jpg\",\n        \"overview\": \"A sudden loss disrupts Carol’s orderly life, propelling her into the dating world for the first time in 20 years. Finally living in the present tense, she finds herself swept up in not one, but two unexpected relationships that challenge her assumptions about what it means to grow old.\",\n        \"fullDuration\": 92,\n        \"firstAirDate\": \"2015-05-15\",\n        \"numberOfSeasons\": 0,\n        \"numberOfEpisodes\": 0,\n        \"originCountry\": [\n            \"US\"\n        ],\n        \"originalLanguage\": \"en\",\n        \"originalTitle\": \"I'll See You in My Dreams\"\n    },\n    \"worstTV\": {\n        \"title\": \"\",\n        \"type\": \"\",\n        \"genre\": null,\n        \"rating\": 0,\n        \"watchedTime\": 0,\n        \"imageUrl\": \"\",\n        \"overview\": \"\",\n        \"fullDuration\": 0,\n        \"firstAirDate\": \"\",\n        \"numberOfSeasons\": 0,\n        \"numberOfEpisodes\": 0,\n        \"originCountry\": null,\n        \"originalLanguage\": \"\",\n        \"originalTitle\": \"\"\n    },\n    \"genres\": [\n        {\n            \"name\": \"Drama\",\n            \"numberOfProductions\": 1,\n            \"timeSpent\": 90\n        },\n        {\n            \"name\": \"Comedy\",\n            \"numberOfProductions\": 1,\n            \"timeSpent\": 90\n        }\n    ],\n    \"bingeSessions\": null,\n    \"trends\": [\n        {\n            \"month\": \"2020.1\",\n            \"timeSpent\": 90\n        }\n    ],\n    \"watchedMovies\": [\n        {\n            \"title\": \"I'll See You in My Dreams\",\n            \"type\": \"Movie\",\n            \"genre\": [\n                \"Drama\",\n                \"Comedy\"\n            ],\n            \"rating\": 6.293,\n            \"watchedTime\": 90\n        }\n    ],\n    \"watchedTV\": null\n}"

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), expectedResult)
}
