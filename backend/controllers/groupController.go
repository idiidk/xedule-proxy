package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idiidk/xedule-proxy/xedule"
	"github.com/idiidk/xedule-proxy/xedule/models"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type GroupController struct{}

func (c GroupController) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("/", func(ctx *gin.Context) {
		groups, err := xedule.GetGroups()

		if err != nil {
			ctx.Error(err)
			return
		}

		filter := ctx.Query("filter")
		if filter != "" {
			var filteredGroups []models.XeduleGroup

			for _, g := range *groups {
				if fuzzy.Match(filter, g.Code) {
					filteredGroups = append(filteredGroups, g)
				}
			}

			ctx.JSON(http.StatusOK, filteredGroups)
			return
		}

		ctx.JSON(http.StatusOK, groups)
	})
}
