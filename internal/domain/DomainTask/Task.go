package DomainTask

import "time"

type Task struct {
	ID          uint      `json:"id" example:"123"`
	Title       string    `json:"title" example:"Hacer tarea"`
	Description string    `json:"description" example:""`
	StartDate   time.Time `json:"startDate" example:"2021-02-24"`
	DueDate     time.Time `json:"dueDate" example:"2021-02-24"`
	Alarm       bool      `json:"alarm" example:"true"`
	Priority    string    `json:"priority" example:"Low"`
	Status      string    `json:"status" example:"Not Started"`
	CreatedAt   time.Time `json:"created_at", omitempty example:"2021-02-24 20:19:39"`
	UpdateAt    time.Time `json:"update_at", omitempty example:"2021-02-24 20:19:39"`
	ListRefer   uint      `json:"listID" example:"123"`
}

type CreateTask struct {
	Title       string `json:"title" example:"Hacer tarea"`
	Description string `json:"description" example:""`
	StartDate   string `json:"startDate" example:"2021-02-24"`
	DueDate     string `json:"dueDate" example:"2021-02-24"`
	Alarm       bool   `json:"alarm" example:"true"`
	Priority    string `json:"priority" example:"Low"`
	Status      string `json:"status" example:"Not Started"`
	ListRefer   uint   `json:"listID" example:"123"`
}

type ReadTask struct {
	Titulo      string `json:"Titulo"`
	Description string `json:"Description"`
}

func ToDomainMapper(createTask *CreateTask) *Task {
	starTime, _ := time.Parse("2006-01-02", createTask.StartDate)
	dueTime, _ := time.Parse("2006-01-02", createTask.DueDate)
	return &Task{
		Title:       createTask.Title,
		Description: createTask.Description,
		StartDate:   starTime,
		DueDate:     dueTime,
		Alarm:       createTask.Alarm,
		Priority:    createTask.Priority,
		Status:      createTask.Status,
		ListRefer:   createTask.ListRefer,
	}
}

func ArrayToDomainMapper(tasks *[]CreateTask) *[]Task {
	domainTask := make([]Task, len(*tasks))
	for i, task := range *tasks {
		domainTask[i] = *ToDomainMapper(&task)
	}
	return &domainTask
}
