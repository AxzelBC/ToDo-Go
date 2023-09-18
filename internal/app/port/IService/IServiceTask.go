package IService

import "github.com/AxzelBC/ToDo-Go/internal/domain"

type ServiceTask interface {
	Create(domain.Task) error
	GetById(id uint) domain.Task
	GetAll() error
}
