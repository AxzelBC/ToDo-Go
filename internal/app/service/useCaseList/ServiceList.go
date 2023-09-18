package usecaselist

import (
	"fmt"

	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IService"
	"github.com/AxzelBC/ToDo-Go/internal/domain"
)

type service struct {
	repository IRepository.IListPort
}

func NewListService(respository IRepository.IListPort) IService.ServiceList {
	return &service{
		repository: respository,
	}
}

func (s *service) Create(domain.List) error {
	return s.repository.Create()
}

func (s *service) GetById(id uint) domain.List {
	resp := s.repository.GetById()
	fmt.Println(resp)
	return domain.List{}
}

func (s *service) GetAll() error {
	return s.repository.GetAll()
}
