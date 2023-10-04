package model

import "time"

type TagModel struct {
	ID        uint      `json:"id" example:"123" gorm:"primaryKey"`
	Title     string    `json:"title" example:"University" gorm:"unique;not null"`
	Color     string    `josn:"color" example:"Red"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
}
