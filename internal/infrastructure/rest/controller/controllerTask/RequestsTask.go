package controllerTask

type NewTaskRequest struct {
	Title       string `json:"title" example:"Hacer tarea"`
	Description string `json:"description" example:""`
	StartDate   string `json:"startDate" example:"2021-02-24"`
	DueDate     string `json:"dueDate" example:"2021-02-24"`
	Alarm       bool   `json:"alarm" example:"true"`
	Priority    string `json:"priority" example:"Low"`
	Status      string `json:"status" example:"Not Started"`
}
