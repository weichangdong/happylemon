package middleware

import (
	"happylemon/lib/log"
	"happylemon/lib/prometheus"

	"github.com/gin-gonic/gin"
)

func MyRecoveryWithWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.ErrorLog(err.(string))
				c.JSON(200, map[string]interface{}{
					"ret": 5,
				})
				c.AbortWithStatus(200)
				prometheus.HttpCodeCount(c, 500)
			}
		}()
		c.Next()
	}
}
