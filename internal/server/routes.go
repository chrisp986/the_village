package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		// v1.GET("/players", getPlayers)
		// v1.GET("/player/:player_id", getPlayer)
		// v1.GET("/active_players", getActivePlayers)
		v1.GET("/", func(c *gin.Context) {
			time.Sleep(2 * time.Second)
			c.String(http.StatusOK, "Welcome Gin Server")
		})
		// v1.POST("/players", postPlayers)
		// v1.POST("/calc_new_res", postCalculateNewResources)
	}
}
