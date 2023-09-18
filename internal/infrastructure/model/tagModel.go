package model

import "gorm.io/gorm"

type TagModel struct {
	Titulo string
	Color  string
	gorm.Model
}
