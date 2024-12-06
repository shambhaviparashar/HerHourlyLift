package routes

import (
	"HerHourlyLiftBackend/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterQuoteRoutes registers routes related to quotes
func RegisterQuoteRoutes(router *gin.Engine) {
	quoteGroup := router.Group("/quotes")
	{
		quoteGroup.GET("/random", handlers.RandomQuoteHandler)
	}
}
