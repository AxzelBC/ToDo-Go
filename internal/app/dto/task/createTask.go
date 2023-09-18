package dto_task

//Id          string `json:"Id"`
type CreateTask struct {
	Titulo      string `json:"Titulo" binding:"required"`
	Description string `json:"Description" binding:"required"`
	FechaInicio string `json:"FechaInicio" binding:"required"`
	FechaFin    string `json:"FechaFin" binding:"required"`
	Alarma      bool   `json:"Alarma" binding:"required"`
	Prioridad   string `json:"Prioridad" binding:"required"`
	Estado      string `json:"Estado" binding:"required"`
}

//Activity    model.ActivityModel `json:"activity"`
