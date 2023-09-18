package model

import (
	"time"

	"gorm.io/gorm"
)

type TaskModel struct {
	Titulo      string
	Description string
	FechaInicio time.Time
	FechaFin    time.Time
	Alarma      bool
	Prioridad   string
	Estado      string
	gorm.Model
}
