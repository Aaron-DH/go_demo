package middleware

import (
	"demo1_gogin_api/utils"
	"demo1_gogin_api/errno"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("X-Auth-Token")

		if len(tokenHeader) == 0 {
			utils.SendResponse(c, errno.ErrTokenExpired, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
