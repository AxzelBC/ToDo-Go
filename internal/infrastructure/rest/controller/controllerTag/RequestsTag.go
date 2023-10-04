package controllerTag

type NewTagRequest struct {
	Title string `json:"title" example:"Favourites" gorm:"unique"`
	Color string `json:"color" example:"Red"`
}
