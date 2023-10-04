package model

import (
	"time"

	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainTask"
)

type TaskModel struct {
	ID          uint       `json:"id" example:"123" gorm:"primaryKey"`
	Title       string     `json:"title" example:"University" gorm:"unique; not null"`
	Description string     `json:"description" example:"This is a example description"`
	StartDate   time.Time  `json:"startDate" example:"2021-02-24" gorm:"not null"`
	DueDate     time.Time  `json:"dueDate" example:"2021-02-24" gorm:"not null"`
	Alarm       bool       `json:"alarm" example:"true" gorm:"default:true"`
	Priority    string     `json:"priority" example:"Low" gorm:"default:Low"`
	Status      string     `json:"status" example:"Not Started" gorm:"default:Not Started"`
	CreatedAt   time.Time  `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	ListModelID uint       `json:"listID" example:"123"`
	ListModel   *ListModel `json:"List" example:"List{}"`
}

func (task *TaskModel) ToDomainMapper() *DomainTask.Task {
	return &DomainTask.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		StartDate:   task.StartDate,
		DueDate:     task.DueDate,
		Alarm:       task.Alarm,
		Priority:    task.Priority,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdateAt:    task.UpdatedAt,
		ListRefer:   task.ListModelID,
	}
}
