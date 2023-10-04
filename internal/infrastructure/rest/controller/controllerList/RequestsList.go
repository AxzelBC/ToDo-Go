package controllerList

import "github.com/AxzelBC/ToDo-Go/internal/domain/DomainTask"

type NewListRequest struct {
	Title  string                  `json:"title" example:"university"`
	Status string                  `json:"status" example:"done"`
	Task   []DomainTask.CreateTask `json:"Tasks" example:"[Task]"`
}
