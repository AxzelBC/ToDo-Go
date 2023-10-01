package model

import (
	"time"

	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainList"
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainTask"
)

type ListModel struct {
	ID        uint        `json:"id" example:"123" gorm:"primaryKey" gorm:"not null"`
	Title     string      `json:"title" example:"University" gorm:"unique"`
	Status    string      `json:"status" example:"Not Started"`
	CreatedAt time.Time   `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time   `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	Tasks     []TaskModel `json:"tasks" example:"[{},{}]"`
}

func (list *ListModel) ToDomainMapper() *DomainList.List {

	taskDomain := make([]DomainTask.Task, len(list.Tasks))
	for i, task := range list.Tasks {
		taskDomain[i] = *task.ToDomainMapper()
	}

	return &DomainList.List{
		ID:        list.ID,
		Title:     list.Title,
		Status:    list.Status,
		Tasks:     taskDomain,
		CreatedAt: list.CreatedAt,
		UpdateAt:  list.UpdatedAt,
	}
}
