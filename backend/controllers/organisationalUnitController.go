package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idiidk/xedule-proxy/xedule"
)

type OrganisationalUnitController struct{}

func (c OrganisationalUnitController) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("/", func(ctx *gin.Context) {
		organisationalUnits, err := xedule.GetOrganisationalUnits()
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, organisationalUnits)
	})
}
