package RepositoryTag

import (
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainTag"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/model"
)

func toDomainMapper(tag *model.TagModel) *DomainTag.Tag {
	return &DomainTag.Tag{
		ID:        uint(tag.ID),
		Title:     tag.Title,
		Color:     tag.Color,
		CreatedAt: tag.CreatedAt,
		UpdateAt:  tag.UpdatedAt,
	}
}

func fromDomainMapper(tag *DomainTag.Tag) *model.TagModel {
	return &model.TagModel{
		ID:        tag.ID,
		Color:     tag.Color,
		Title:     tag.Title,
		CreatedAt: tag.CreatedAt,
		UpdatedAt: tag.UpdateAt,
	}
}

func arrayToDomainMapper(tags *[]model.TagModel) *[]DomainTag.Tag {
	tagDomain := make([]DomainTag.Tag, len(*tags))
	for i, tag := range *tags {
		tagDomain[i] = *toDomainMapper(&tag)
	}
	return &tagDomain
}
