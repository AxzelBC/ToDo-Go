package controllerList

type ResponseNewList struct {
	ID     uint   `json:"id" example:"123"`
	Title  string `json:"title" example:"University"`
	Status string `json:"status" example:"Not Started"`
}
