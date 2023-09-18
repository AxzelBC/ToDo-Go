package routes

import (
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/controller/controllerList"
	"github.com/gin-gonic/gin"
)

func RoutesList(router *gin.RouterGroup, controller *controllerList.Controller) {
	routerList := router.Group("/list")
	{
		routerList.POST("/", controller.NewList)
		routerList.GET("/:id", controller.GetListByID)
		routerList.GET("/", controller.GetAllLists)
	}
}
