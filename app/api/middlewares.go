package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddlewares(r *gin.Engine) {
	// set start time on every request to check response time
	r.Use(func(c *gin.Context) {
		startTime := time.Now()
		c.Set("startTime", startTime)
		c.Next()
	})
}
