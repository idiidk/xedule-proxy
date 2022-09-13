package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idiidk/xedule-proxy/xedule"
)

type GroupController struct{}

func (c GroupController) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("/", func(ctx *gin.Context) {
		groups, err := xedule.GetGroups()
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, groups)
	})
}
