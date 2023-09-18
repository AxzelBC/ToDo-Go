package domain

import "time"

type Task struct {
	ID          uint      `json:"id" example:"123"`
	Titulo      string    `json:"titulo" example:"Hacer tarea"`
	Description string    `json:"description" example:""`
	FechaInicio time.Time `json:"fecha_inicio" example:"2021-02-24"`
	FechaFin    time.Time `json:"fecha_fin" example:"2021-02-24"`
	Alarma      bool      `json:"alarma" example:""`
	Prioridad   string    `json:"prioridad" example:""`
	Estado      string    `json:"estado" example:""`
	CreatedAt   time.Time `json:"created_at", omitempty example:"2021-02-24 20:19:39"`
	UpdateAt    time.Time `json:"update_at", omitempty example:"2021-02-24 20:19:39"`
}
