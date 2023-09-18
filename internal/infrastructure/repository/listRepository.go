package repository

import (
	"fmt"

	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/model"
)

const (
	sin_empezar string = "Sin Empezar"
	en_progreso string = "En Progreso"
	completa    string = "Completa"
)

type listRepository struct {
	db IRepository.IDatabasePort
}

func NewIListPort(db IRepository.IDatabasePort) IRepository.IListPort {
	db.GetDB().AutoMigrate(&model.ListModel{})
	return &listRepository{
		db: db,
	}
}

func (lr *listRepository) Create() error {
	created := lr.db.GetDB().Create(&model.ListModel{
		Titulo: "Universidad",
		Estado: sin_empezar,
	})

	if created.Error != nil {
		fmt.Printf("%+v\n", created.Error)
		return created.Error
	}

	fmt.Printf("%+v\n", created)

	return nil
}

func (lr *listRepository) GetAll() error {

	var lists []model.ListModel

	read := lr.db.GetDB().Find(&lists)

	if read.Error != nil {
		fmt.Printf("%+v\n", read.Error)
		return read.Error
	}

	fmt.Printf("%+v\n", lists)

	return nil
}

func (lr *listRepository) GetById() error {

	read := model.ListModel{}

	err := lr.db.GetDB().First(&read, "ID = ?", "1")

	if err.Error != nil {
		fmt.Printf("%+v\n", err.Error)
		return err.Error
	}

	fmt.Printf("%+v\n", read)

	return nil
}

/*
func (lr *listRepository) UpdateList() error {
	lr.db.GetDB()
}
func (lr *listRepository) DeleteList() error {
	lr.db.GetDB()
}
*/
