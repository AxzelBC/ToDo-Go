package routes

import (
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/controller/controllerTask"
	"github.com/gin-gonic/gin"
)

func RoutesTask(router *gin.RouterGroup, controller *controllerTask.Controller) {
	routerTask := router.Group("/task/:listID")
	{
		routerTask.POST("/", controller.NewTask)
		routerTask.GET("/:id", controller.GetTaskById)
		routerTask.GET("/", controller.GetAllTaks)
	}
}
