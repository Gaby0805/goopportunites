package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func inicializateRouter(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK,gin.H{
				"msg": "get opening",
			})
		})
	
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK,gin.H{
				"msg": "Post opening",
			})
		})
	
		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK,gin.H{
				"msg": "DELETE opening",
			})
		})
	
		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK,gin.H{
				"msg": "PUT opening",
			})
		})
	
		v1.GET("/openingS", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK,gin.H{
				"msg": "get opening",
			})
		})
	}

}
