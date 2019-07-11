package router

import (
	"github.com/gin-gonic/gin"
	"net/http"

	_ "demo1_gogin_api/docs"
	"demo1_gogin_api/router/middleware"
	"demo1_gogin_api/service"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.POST("/login", service.Login)

	checkservice := g.Group("/check")
	{
		checkservice.GET("/health", service.HealthCheck)
	}

	tagservice := g.Group("/v1/tags")
	tagservice.Use(middleware.AuthMiddleware())
	{
		tagservice.GET("", service.GetTags)
		tagservice.POST("", service.CreateTag)
		tagservice.PUT("/:id", service.UpdateTag)
	}

	userservice := g.Group("/v1/users")
	userservice.Use(middleware.AuthMiddleware())
	{
		userservice.GET("", service.GetUsers)
	}

	return g
}
