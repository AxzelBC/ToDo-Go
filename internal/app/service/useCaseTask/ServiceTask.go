package usecasetask

import (
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IService"
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainTask"
)

type service struct {
	repository IRepository.ITaskPort
}

func NewTaskService(repository IRepository.ITaskPort) IService.ServiceTask {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(createTask *DomainTask.CreateTask) (*DomainTask.Task, error) {
	taskModel := DomainTask.ToDomainMapper(createTask)
	return s.repository.Create(taskModel)
}

func (s *service) GetById(id uint) (*DomainTask.Task, error) {
	return s.repository.GetById(id)
}

func (s *service) GetAll(id uint) (*[]DomainTask.Task, error) {
	return s.repository.GetAll(id)
}
