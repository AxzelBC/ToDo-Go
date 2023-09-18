package repository

import (
	"fmt"

	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/model"
)

type tagRepository struct {
	db IRepository.IDatabasePort
}

func NewITagPort(db IRepository.IDatabasePort) IRepository.ITagPort {
	db.GetDB().AutoMigrate(&model.TagModel{})
	return &tagRepository{
		db: db,
	}
}

// Create a Tag
func (tr *tagRepository) Create() error {
	created := tr.db.GetDB().Create(&model.TagModel{
		Titulo: "Bug",
		Color:  "azul",
	})

	fmt.Printf("%+v", created)

	return nil
}

// List a Tag
func (tr *tagRepository) GetAll() error {
	var tags []model.TagModel

	read := tr.db.GetDB().Find(tags)

	if read.Error != nil {
		return read.Error
	}

	fmt.Printf("%+v", tags)

	return nil
}

func (tr *tagRepository) GetById() error {
	var tag model.TagModel

	read := tr.db.GetDB().First(&tag, "ID = ?", "1")

	if read.Error != nil {
		return read.Error
	}

	fmt.Printf("%+v", tag)

	return nil
}
