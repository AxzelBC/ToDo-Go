package IService

import "github.com/AxzelBC/ToDo-Go/internal/domain"

type ServiceTag interface {
	Create(domain.Tag) error
	GetById(id uint) domain.Tag
	GetAll() error
}
