package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/idiidk/xedule-proxy/controllers"
)

func InitRoutes(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		statusGroup := v1.Group("/status")
		{
			statusController := new(controllers.StatusController)
			statusController.RegisterRoutes(statusGroup)
		}
	}
}
