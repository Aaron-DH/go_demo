package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
        c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
        c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
        c.Next()
}

func Options(c *gin.Context) {
	if c.Request.Method != "Options" {
		c.Next()
	} else {
                c.Header("Access-Control-Allow-Origin", "*")
                c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
                c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
                c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
                c.Header("Content-Type", "application/json")
                c.AbortWithStatus(200)
	}
}
