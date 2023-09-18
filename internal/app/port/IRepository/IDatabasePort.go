package IRepository

import (
	"io"

	"gorm.io/gorm"
)

type IDatabasePort interface {
	GetDB() *gorm.DB
	io.Closer
}
