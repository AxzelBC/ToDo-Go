package domain

import "time"

type List struct {
	ID        uint      `json:"id" example:"123"`
	Titulo    string    `json:"titulo" example:"university"`
	Estado    string    `json:"estado" example:"done"`
	CreatedAt time.Time `json:"created_at", omitempty example:"2021-02-24 20:19:39"`
	UpdateAt  time.Time `json:"update_at", omitempty example:"2021-02-24 20:19:39"`
}
