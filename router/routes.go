package router

import (
	"github.com/franciscojoaopedro/goopportunities.git/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpiningHandler)
		v1.GET("/openings", handler.ListAllOpiningHandler)
		v1.POST("/opening", handler.CreateOpiningHandler)
		v1.PUT("/opening", handler.UpdateOpiningHandler)
		v1.DELETE("/opening", handler.DeleteOpiningHandler)
	}

}
