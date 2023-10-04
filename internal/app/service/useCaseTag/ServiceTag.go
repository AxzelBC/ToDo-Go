package usecasetag

import (
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IService"
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainTag"
)

type service struct {
	repository IRepository.ITagPort
}

func NewTagService(respository IRepository.ITagPort) IService.ServiceTag {
	return &service{
		repository: respository,
	}
}

func (s *service) Create(createTag *DomainTag.CreateTag) (*DomainTag.Tag, error) {
	tagModel := DomainTag.ToDomainMapper(createTag)
	return s.repository.Create(tagModel)
}
func (s *service) GetById(id uint) (*DomainTag.Tag, error) {
	return s.repository.GetById(id)
}
func (s *service) GetAll() (*[]DomainTag.Tag, error) {
	return s.repository.GetAll()
}
