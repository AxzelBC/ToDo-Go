package IService

import (
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainList"
)

type ServiceList interface {
	Create(*DomainList.CreateList) (*DomainList.List, error)
	GetById(idList uint) (*DomainList.List, error)
	GetAll() (*[]DomainList.List, error)
}
