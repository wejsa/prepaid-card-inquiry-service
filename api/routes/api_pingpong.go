package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/serverkona1/prepaid-card-inquiry-service/service/pingpong"
)

func initPingPongApi(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		message := pingpong.SayPongToPing()
		c.JSON(200, gin.H{"message": message})
	})
	return r
}
