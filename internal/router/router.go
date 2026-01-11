package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msterhuj/ratioarr/internal/static"
	"github.com/msterhuj/ratioarr/internal/views"
)

func NewRouter() *gin.Engine {
	// Gin mode (release/debug will be set later via log_level)
	r := gin.New()

	// Global middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// Disable trusted proxy warning.
	r.SetTrustedProxies(nil)

	// Healthcheck
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &views.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", views.Index())
	})

	// add static path and read from embedded static
	r.StaticFS("/static", http.FS(static.Files))

	// API group
	/*api := r.Group("/api")
	{
		api.GET("/trackers", listTrackers)
		api.GET("/ratios", listRatios)
	}
*/
	return r
}
