package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postData(c *gin.Context) {
	var newEntity Data

	if err := c.BindJSON(&newEntity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ERROR: Sent data is incorrect (check if your JSON file is in the correct format)",
		})
		return
	}

	report, err := generateReport(&newEntity)

	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "Empty history" || err.Error() == "Empty data" || err.Error() == "Not enough data" {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{
			"error": "ERROR: " + err.Error(),
		})
	} else {
		c.IndentedJSON(http.StatusOK, report)
	}
}
