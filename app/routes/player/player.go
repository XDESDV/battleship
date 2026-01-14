package player

import (
	controller "battleship/app/controllers/player"
	service "battleship/app/services/player"

	"github.com/gin-gonic/gin"
)

func SetupRouter(g *gin.Engine) {

	servicesPlayer := service.New()
	playerController := controller.New(servicesPlayer)

	v1 := g.Group("/v1")
	{
		players := v1.Group("/players")
		{
			players.POST("", playerController.Create)
			players.GET("", playerController.Get)
			players.GET("/:id", playerController.GetByID)
			players.POST("/:id", playerController.Update)
			players.POST("/:id/suspend", playerController.Suspend)
			players.GET("/IDS/:ids", playerController.GetByIDs)
		}
	}
}
