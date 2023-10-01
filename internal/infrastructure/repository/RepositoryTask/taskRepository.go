package RepositoryTask

import (
	"log"

	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainTask"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/model"
)

type taskRepository struct {
	db IRepository.IDatabasePort
}

func NewITaskPort(db IRepository.IDatabasePort) IRepository.ITaskPort {
	return &taskRepository{
		db: db,
	}
}

func (tr *taskRepository) Create(createTask *DomainTask.Task) (*DomainTask.Task, error) {
	taskModel := FromDomainMapper(createTask)
	var listFetched model.ListModel

	// find list
	err := tr.db.GetDB().Debug().First(&listFetched, createTask.ListRefer).Error

	if err != nil {
		log.Printf("%+v\n", err)
		return &DomainTask.Task{}, err
	}

	taskModel.ListModel = &listFetched

	err = tr.db.GetDB().Debug().Create(&taskModel).Error

	if err != nil {
		log.Printf("%+v\n", err)
		return &DomainTask.Task{}, err
	}

	return taskModel.ToDomainMapper(), nil
}

func (tr *taskRepository) GetAll(idList uint) (*[]DomainTask.Task, error) {

	var tasks []model.TaskModel

	err := tr.db.GetDB().Debug().Find(&tasks).Where("ListModelID = ?", idList).Error

	if err != nil {
		var errArray *[]DomainTask.Task
		log.Printf("%+v\n", err)
		return errArray, err
	}

	return ArrayToDomainMapper(&tasks), nil
}

func (tr *taskRepository) GetById(idTask uint) (*DomainTask.Task, error) {

	var task model.TaskModel

	err := tr.db.GetDB().First(&task, "ID = ?", idTask).Error

	if err != nil {
		log.Printf("%+v\n", err)
		return &DomainTask.Task{}, err
	}

	return task.ToDomainMapper(), nil
}
