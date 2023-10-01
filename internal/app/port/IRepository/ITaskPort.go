package IRepository

import "github.com/AxzelBC/ToDo-Go/internal/domain/DomainTask"

type ITaskPort interface {
	Create(*DomainTask.Task) (*DomainTask.Task, error)
	GetById(uint) (*DomainTask.Task, error)
	GetAll(uint) (*[]DomainTask.Task, error)
}
