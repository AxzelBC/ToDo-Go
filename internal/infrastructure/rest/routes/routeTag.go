package routes

import (
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/controller/controllerTag"
	"github.com/gin-gonic/gin"
)

func RoutesTag(router *gin.RouterGroup, controller *controllerTag.Controller) {
	routerTag := router.Group("/tag")
	{
		routerTag.POST("/", controller.NewTag)
		routerTag.GET("/:id", controller.GetTagByID)
		routerTag.GET("/", controller.GetAllTags)
	}
}
