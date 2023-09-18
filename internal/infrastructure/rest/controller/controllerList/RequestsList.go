package controllerList

type NewListRequest struct {
	Titulo string `json:"titulo" example:"university" gorm:"unique" binding:"requierd"`
	Estado string `json:"estado" example:"done" binding:"requierd"`
}
