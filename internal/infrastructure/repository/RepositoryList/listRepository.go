package RepositoryList

import (
	"fmt"
	"log"

	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainList"
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainTask"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/model"
)

const (
	not_started string = "Not Started"
	in_progress string = "In Progress"
	done        string = "Done"
)

type listRepository struct {
	db IRepository.IDatabasePort
}

func NewIListPort(db IRepository.IDatabasePort) IRepository.IListPort {
	return &listRepository{
		db: db,
	}
}

func (lr *listRepository) Create(newList *DomainList.List) (*DomainList.List, error) {

	// become domain to model
	listFetched, tasksFetched := FromDomainMapper(newList)

	// create list
	err := lr.db.GetDB().Debug().Create(&listFetched).Error

	// if exist a error.
	if err != nil {
		return &DomainList.List{}, err
	}

	// if the tasks is nil
	if tasksFetched == nil {
		listCreated := listFetched.ToDomainMapper()
		return listCreated, nil
	}

	tasksToDomain := make([]DomainTask.Task, len(*tasksFetched))

	// create tasks associated with the list
	for i, task := range *tasksFetched {

		task.ListModel = listFetched

		err = lr.db.GetDB().Create(&task).Error

		if err != nil {
			return &DomainList.List{}, err
		}

		tasksToDomain[i] = *task.ToDomainMapper()
	}

	// Mapping model to domain
	createdList := listFetched.ToDomainMapper()

	// add tasks to the List
	createdList.Tasks = tasksToDomain

	return createdList, nil
}

func (lr *listRepository) GetAll() (*[]DomainList.List, error) {

	var lists []model.ListModel
	err := lr.db.GetDB().Find(&lists).Error

	if err != nil {
		fmt.Printf("%+v\n", err)
		return nil, err
	}

	return ArrayToDomainMapper(&lists), nil
}

func (lr *listRepository) GetById(idList uint) (*DomainList.List, error) {

	var listFetched model.ListModel

	err := lr.db.GetDB().Debug().Preload("Tasks").Find(&listFetched, idList).Error

	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return listFetched.ToDomainMapper(), nil
}
