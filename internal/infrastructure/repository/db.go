package repository

import (
	"github.com/AxzelBC/ToDo-Go/internal/app/port/IRepository"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/config"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type database struct {
	*gorm.DB
}

func NewDB(config *config.DatabaseConfig) (IRepository.IDatabasePort, error) {
	db, err := newDatabase(config)

	if err != nil {
		return nil, err
	}

	return &database{
		db,
	}, nil
}

func newDatabase(conf *config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.sql"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.TagModel{}, &model.ListModel{}, &model.TaskModel{})

	if err != nil {
		return nil, err
	}

	return db, err
}

func (db *database) GetDB() *gorm.DB {
	return db.DB
}

func (db *database) Close() error {
	return db.Close()
}
