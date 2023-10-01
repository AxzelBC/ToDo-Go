package usecaselist

import (
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IService"
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainList"
)

type service struct {
	repository IRepository.IListPort
}

func NewListService(respository IRepository.IListPort) IService.ServiceList {
	return &service{
		repository: respository,
	}
}

func (s *service) Create(newList *DomainList.CreateList) (*DomainList.List, error) {
	listCreate := DomainList.ToDomainMapper(newList)
	return s.repository.Create(listCreate)
}

func (s *service) GetById(idList uint) (*DomainList.List, error) {
	return s.repository.GetById(idList)
}

func (s *service) GetAll() (*[]DomainList.List, error) {
	return s.repository.GetAll()
}
