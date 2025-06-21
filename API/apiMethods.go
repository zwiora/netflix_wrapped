package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postData(c *gin.Context) {
	var newEntity Data

	if err := c.BindJSON(&newEntity); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Sent data is incorrect (check if your JSON file is in the correct format)",
		})
	}

	report, err := generateReport(&newEntity)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.IndentedJSON(http.StatusOK, report)
	}
}
