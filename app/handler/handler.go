package handler

import (
	_ "embed"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed colours.json
var coloursJSON []byte

func Favicon(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func GetLanguages(c *gin.Context) {
	// Parse JSON data from embedded file
	var languages map[string]string
	err := json.Unmarshal(coloursJSON, &languages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to parse JSON data",
		})
		return
	}

	// Return all language entries
	c.JSON(http.StatusOK, gin.H{
		"languages": languages,
		"count":     len(languages),
	})
}
