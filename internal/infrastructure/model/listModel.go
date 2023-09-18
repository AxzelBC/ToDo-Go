package model

import "gorm.io/gorm"

type ListModel struct {
	Titulo string
	Estado string
	gorm.Model
}
