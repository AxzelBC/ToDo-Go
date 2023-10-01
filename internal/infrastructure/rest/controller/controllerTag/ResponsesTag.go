package controllerTag

type ResponseNewTag struct {
	ID    uint   `json:"id" example:"123"`
	Title string `json:"title" example:"Favourites"`
	Color string `json:"color" example:"Red"`
}
