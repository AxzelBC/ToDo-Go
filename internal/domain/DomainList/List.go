package DomainList

import (
	"time"

	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainTask"
)

type List struct {
	ID        uint              `json:"id" example:"123"`
	Title     string            `json:"title" example:"university"`
	Status    string            `json:"status" example:"done"`
	CreatedAt time.Time         `json:"created_at", omitempty example:"2021-02-24 20:19:39"`
	UpdateAt  time.Time         `json:"update_at", omitempty example:"2021-02-24 20:19:39"`
	Tasks     []DomainTask.Task `json:"tasks" example:"[{},{}]"`
}

type CreateList struct {
	Title  string                  `json:"title" example:"University"`
	Status string                  `json:"status" example:"Not Started"`
	Tasks  []DomainTask.CreateTask `json:"tasks" example:"[Task:{}]"`
}

func ToDomainMapper(createList *CreateList) *List {
	test := *DomainTask.ArrayToDomainMapper(&createList.Tasks)
	return &List{
		Title:  createList.Title,
		Status: createList.Status,
		Tasks:  test,
	}
}
