package IRepository

import "github.com/AxzelBC/ToDo-Go/internal/domain/DomainTag"

type ITagPort interface {
	Create(*DomainTag.Tag) (*DomainTag.Tag, error)
	GetById(uint) (*DomainTag.Tag, error)
	GetAll() (*[]DomainTag.Tag, error)
}
