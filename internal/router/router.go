package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// Gin mode (release/debug will be set later via log_level)
	r := gin.New()

	// Global middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Healthcheck
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API group
	/*api := r.Group("/api")
	{
		api.GET("/trackers", listTrackers)
		api.GET("/ratios", listRatios)
	}
*/
	return r
}
