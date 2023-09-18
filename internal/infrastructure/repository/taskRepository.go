package repository

import (
	"fmt"
	"time"

	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/model"
)

type taskRepository struct {
	db IRepository.IDatabasePort
}

func NewITaskPort(db IRepository.IDatabasePort) IRepository.ITaskPort {
	db.GetDB().AutoMigrate(&model.TaskModel{})
	return &taskRepository{
		db: db,
	}
}

func (tr *taskRepository) Create() error {

	// string to time.time
	initTime, _ := time.Parse("2006-01-02", "2023-09-11")
	finalTime, _ := time.Parse("2006-01-02", "2023-09-12")

	created := tr.db.GetDB().Create(&model.TaskModel{
		Titulo:      "Hacer ToDo-Go",
		Description: "Se debe realizar un api-rest en Golang sobre una ToDo-List",
		FechaInicio: initTime,
		FechaFin:    finalTime,
		Alarma:      true,
		Prioridad:   "Alta",
		Estado:      sin_empezar,
	})

	if created.Error != nil {
		fmt.Printf("%+v\n", created.Error)
		return created.Error
	}

	fmt.Printf("%+v\n", created)

	return nil
}

func (tr *taskRepository) GetAll() error {

	var tasks []model.TagModel

	read := tr.db.GetDB().Find(tasks)

	if read.Error != nil {
		fmt.Printf("%+v\n", read.Error)
		return read.Error
	}

	return nil
}

func (tr *taskRepository) GetById() error {

	var task model.TagModel

	read := tr.db.GetDB().First(&task, "ID = ?", "1")

	if read.Error != nil {
		fmt.Printf("%+v\n", read.Error)
		return read.Error
	}

	fmt.Printf("%+v\n", task)

	return nil
}
