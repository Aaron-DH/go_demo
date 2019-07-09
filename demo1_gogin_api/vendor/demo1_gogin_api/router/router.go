package router

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"demo1_gogin_api/service"
	"demo1_gogin_api/router/middleware"
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
                checkservice.GET("/health", service.HealthCheck)
        }

        tagservice := g.Group("/tag")
		{
			tagservice.GET("/tags", service.GetTags)
			tagservice.POST("/tags", service.CreateTag)
			tagservice.PUT("/tags/:id", service.UpdateTag)
		}

		userservice := g.Group("/user")
		{
			userservice.GET("/users", service.GetUsers)
			userservice.POST("/login", service.Login)
		}

        return g
}
