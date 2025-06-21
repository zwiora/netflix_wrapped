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

// MOCK: podmieniamy generateReport na wersję testową
var mockGenerateReport func(data *Data) (Report, error)

func init() {
	// External API
	var err error
	setApiKey()
	tmdbClient, err = initializeTMDB()
	if err != nil {
		log.Println(err)
	}

	// Podmieniamy globalną funkcję generateReport na mock
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
	assert.Contains(t, w.Body.String(), "Not enough data")
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

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User1")
}
