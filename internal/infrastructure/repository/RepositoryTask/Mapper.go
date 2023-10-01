package RepositoryTask

import (
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainTask"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/model"
)

func FromDomainMapper(task *DomainTask.Task) *model.TaskModel {
	return &model.TaskModel{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		StartDate:   task.StartDate,
		DueDate:     task.DueDate,
		Alarm:       task.Alarm,
		Priority:    task.Priority,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdateAt,
		ListModelID: task.ListRefer,
	}
}

func ArrayToDomainMapper(tasks *[]model.TaskModel) *[]DomainTask.Task {
	taskDomain := make([]DomainTask.Task, len(*tasks))
	for i, task := range *tasks {
		taskDomain[i] = *task.ToDomainMapper()
	}
	return &taskDomain
}

func ArrayFromDomainMapper(tasks *[]DomainTask.Task) *[]model.TaskModel {
	taskDomain := make([]model.TaskModel, len(*tasks))
	for i, task := range *tasks {
		taskDomain[i] = *(FromDomainMapper(&task))
	}
	return &taskDomain
}
