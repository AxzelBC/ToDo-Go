package adapterList

import (
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	usecaselist "github.com/AxzelBC/ToDo-Go/internal/app/service/useCaseList"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/repository/RepositoryList"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/controller/controllerList"
)

func AdapterList(db *IRepository.IDatabasePort) *controllerList.Controller {
	repository := RepositoryList.NewIListPort(*db)
	service := usecaselist.NewListService(repository)
	return &controllerList.Controller{ServiceList: service}
}
