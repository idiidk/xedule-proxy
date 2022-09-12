package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusController struct{}

func (c StatusController) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "running"})
	})
}
