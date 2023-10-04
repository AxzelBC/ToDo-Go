package RepositoryList

import (
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainList"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/model"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/repository/RepositoryTask"
)

func FromDomainMapper(list *DomainList.List) (*model.ListModel, *[]model.TaskModel) {
	tasksModel := *RepositoryTask.ArrayFromDomainMapper(&list.Tasks)
	return &model.ListModel{
		ID:        list.ID,
		Status:    list.Status,
		Title:     list.Title,
		Tasks:     make([]model.TaskModel, 0),
		CreatedAt: list.CreatedAt,
		UpdatedAt: list.UpdateAt,
	}, &tasksModel
}

func ArrayToDomainMapper(lists *[]model.ListModel) *[]DomainList.List {
	listDomain := make([]DomainList.List, len(*lists))
	for i, list := range *lists {
		listDomain[i] = *list.ToDomainMapper()
	}
	return &listDomain
}
