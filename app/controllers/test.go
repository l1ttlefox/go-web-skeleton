package controllers

import (
	"github.com/gin-gonic/gin"
	"go-web-skeleton/app/models"
)

func Test(ctx *gin.Context) {
	id := ctx.Param("id")
	u := models.GetUserById(id)
	ctx.JSON(200, gin.H{"data": u})
}
