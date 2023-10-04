package IService

import (
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainTag"
)

type ServiceTag interface {
	Create(*DomainTag.CreateTag) (*DomainTag.Tag, error)
	GetById(id uint) (*DomainTag.Tag, error)
	GetAll() (*[]DomainTag.Tag, error)
}
