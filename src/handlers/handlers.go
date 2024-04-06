package handlers

import (
	"ASTRA/src/utils"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

func IngestData(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error while reading payload",
			"error":   err,
		})
	}

	err = utils.WriteToFile(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error while writing data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Received POST request successfully",
	})
}
