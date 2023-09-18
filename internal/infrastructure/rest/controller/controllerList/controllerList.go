package controllerList

import (
	"net/http"

	"github.com/AxzelBC/ToDo-Go/internal/app/port/IService"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	ServiceList IService.ServiceList
}

func (c *Controller) NewList(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"data": "Created",
	})
	return
}

func (c *Controller) GetListByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Get",
	})
	return
}

func (c *Controller) GetAllLists(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "GetAll",
	})
	return
}
