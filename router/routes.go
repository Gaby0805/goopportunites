package router

import (
	"github.com/Gaby0805/goopportunites/handler"
	"github.com/gin-gonic/gin"
)

func inicializateRouter(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/opening", handler.CreateOpeningHandler)
		v1.POST("/opening", handler.ShowOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)

	}

}
