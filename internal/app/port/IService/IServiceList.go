package IService

import "github.com/AxzelBC/ToDo-Go/internal/domain"

type ServiceList interface {
	Create(domain.List) error
	GetById(id uint) domain.List
	GetAll() error
}
