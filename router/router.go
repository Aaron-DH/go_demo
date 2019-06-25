package router

import (
	"net/http"
	"github.com/gin-gonic/gin"

        "go_api_demo/service/check"
  	"go_api_demo/service/tag"
        "go_api_demo/router/middleware"
)

func Load(g *gin.Engine, mws ...gin.HandlerFunc) *gin.Engine {
        // Middlewares.
        g.Use(gin.Recovery())
        g.Use(middleware.NoCache)
        g.Use(middleware.Options)
        g.Use(mws...)

        g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "API not found.")
	})

        checkservice := g.Group("/check")
        {
                checkservice.GET("/health", check.HealthCheck)
        }

        tagservice := g.Group("/tag")
	{
		tagservice.GET("/tags", tag.GetTags)
	}

        return g
}
