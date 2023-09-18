package routes

import (
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	adapter "github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/adapter/List"
	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine, db *IRepository.IDatabasePort) {
	routes := router.Group("/api")
	{
		RoutesList(routes, adapter.AdapterList(*db))
	}
}
