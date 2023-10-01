package adapterTask

import (
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	usecasetask "github.com/AxzelBC/ToDo-Go/internal/app/service/useCaseTask"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/repository/RepositoryTask"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/controller/controllerTask"
)

func AdapterTask(db *IRepository.IDatabasePort) *controllerTask.Controller {
	repository := RepositoryTask.NewITaskPort(*db)
	service := usecasetask.NewTaskService(repository)
	return &controllerTask.Controller{ServiceTask: service}
}
