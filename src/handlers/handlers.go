package handlers

import (
	db "ASTRA/src/database"
	"net/http"

	"github.com/labstack/echo"
)

func IngestData(c echo.Context) error {
	// Bind the request body to the RequestBody struct
	var requestBody db.DataLine
	if err := c.Bind(&requestBody); err != nil {
		return err
	}
	err := db.Insert(requestBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error while updating database",
			"error":   err,
		})
	}

	// Process the request data
	// In this example, we'll just return the received data as JSON
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Received POST request successfully",
		"city":    requestBody.City,
		"time":    requestBody.Timestamp,
	})
}
