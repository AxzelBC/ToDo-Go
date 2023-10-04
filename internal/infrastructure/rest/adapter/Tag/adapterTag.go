package adapterTag

import (
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	usecasetag "github.com/AxzelBC/ToDo-Go/internal/app/service/useCaseTag"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/repository/RepositoryTag"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/controller/controllerTag"
)

func AdapterTag(db *IRepository.IDatabasePort) *controllerTag.Controller {
	repository := RepositoryTag.NewITagPort(*db)
	service := usecasetag.NewTagService(repository)
	return &controllerTag.Controller{ServiceTag: service}
}
