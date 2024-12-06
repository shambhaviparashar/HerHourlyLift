package handlers

import (
	"net/http"

	"HerHourlyLiftBackend/controllers"

	"github.com/gin-gonic/gin"
)

// RandomQuoteHandler serves a random quote
func RandomQuoteHandler(c *gin.Context) {
	quote, err := controllers.GetRandomQuote()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch random quote"})
		return
	}
	c.JSON(http.StatusOK, quote)
}
