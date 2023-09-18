package usecasetask

import (
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IService"
	"github.com/AxzelBC/ToDo-Go/internal/domain"
)

type service struct {
	repository IRepository.ITaskPort
}

func NewTaskService(repository IRepository.ITaskPort) IService.ServiceTask {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(domain.Task) error {
	return s.repository.Create()
}

func (s *service) GetById(id uint) domain.Task {
	//s.repository.GetById()
	return domain.Task{}
}

func (s *service) GetAll() error {
	return s.repository.GetAll()
}
