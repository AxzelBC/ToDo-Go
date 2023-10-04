package routes

import (
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	adapterList "github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/adapter/List"
	adapterTag "github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/adapter/Tag"
	adapterTask "github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/adapter/Task"
	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine, db *IRepository.IDatabasePort) {
	routes := router.Group("/api")
	{
		RoutesList(routes, adapterList.AdapterList(db))
		RoutesTag(routes, adapterTag.AdapterTag(db))
		RoutesTask(routes, adapterTask.AdapterTask(db))
	}
}
