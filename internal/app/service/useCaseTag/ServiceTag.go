package usecasetag

import (
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IService"
	"github.com/AxzelBC/ToDo-Go/internal/domain"
)

type service struct {
	repository IRepository.ITagPort
}

func NewTagService(respository IRepository.ITagPort) IService.ServiceTag {
	return &service{
		repository: respository,
	}
}

func (s *service) Create(domain.Tag) error {
	return s.repository.Create()
}
func (s *service) GetById(id uint) domain.Tag {
	//s.repository.GetById()
	return domain.Tag{}
}
func (s *service) GetAll() error {
	return s.repository.GetAll()
}
