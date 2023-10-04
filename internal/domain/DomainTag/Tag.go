package DomainTag

import "time"

type Tag struct {
	ID        uint      `json:"id" example:"1"`
	Title     string    `json:"title" example:"import"`
	Color     string    `json:"color" example:"red"`
	CreatedAt time.Time `json:"created_at", omitempty example:"2021-02-24 20:19:39"`
	UpdateAt  time.Time `json:"update_at", omitempty example:"2021-02-24 20:19:39"`
}

type CreateTag struct {
	Title string `json:"title" example:"import"`
	Color string `json:"color" example:"red"`
}

func ToDomainMapper(createTag *CreateTag) *Tag {
	return &Tag{
		Title: createTag.Title,
		Color: createTag.Color,
	}
}
