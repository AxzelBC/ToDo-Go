package IRepository

import (
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainList"
)

type IListPort interface {
	Create(*DomainList.List) (*DomainList.List, error)
	GetAll() (*[]DomainList.List, error)
	GetById(idList uint) (*DomainList.List, error)
}
