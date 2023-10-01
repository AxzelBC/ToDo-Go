package RepositoryTag

import (
	"log"

	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainTag"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/model"
)

type tagRepository struct {
	db IRepository.IDatabasePort
}

func NewITagPort(db IRepository.IDatabasePort) IRepository.ITagPort {
	return &tagRepository{
		db: db,
	}
}

// Create a Tag
func (tr *tagRepository) Create(createTag *DomainTag.Tag) (*DomainTag.Tag, error) {

	tag := fromDomainMapper(createTag)

	err := tr.db.GetDB().Create(tag).Error

	if err != nil {
		log.Printf("\n%+v\n", err)
		return nil, err
	}

	log.Printf("\n%+v\n", err)

	return toDomainMapper(tag), nil
}

// Get a tag by ID
func (tr *tagRepository) GetById(idTag uint) (*DomainTag.Tag, error) {

	var tag model.TagModel

	err := tr.db.GetDB().First(&tag, "ID = ?", idTag).Error

	if err != nil {
		return nil, err
	}

	return toDomainMapper(&tag), nil
}

// List a Tag
func (tr *tagRepository) GetAll() (*[]DomainTag.Tag, error) {

	var tags []model.TagModel

	err := tr.db.GetDB().Find(&tags).Error

	if err != nil {
		return nil, err
	}

	return arrayToDomainMapper(&tags), nil
}
