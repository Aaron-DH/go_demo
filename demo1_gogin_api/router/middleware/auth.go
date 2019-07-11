package middleware

import (
	"demo1_gogin_api/errno"
	"demo1_gogin_api/redis"
	"demo1_gogin_api/utils"
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

		value, err := redis.Get(tokenHeader)
		if err != nil || value == nil {
			utils.SendResponse(c, errno.ErrTokenExpired, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
