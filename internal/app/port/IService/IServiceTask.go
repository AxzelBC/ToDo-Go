package IService

import (
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainTask"
)

type ServiceTask interface {
	Create(*DomainTask.CreateTask) (*DomainTask.Task, error)
	GetById(uint) (*DomainTask.Task, error)
	GetAll(uint) (*[]DomainTask.Task, error)
}
