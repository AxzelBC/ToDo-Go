package controllerTag

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AxzelBC/ToDo-Go/internal/app/port/IService"
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainTag"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/utils"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	ServiceTag IService.ServiceTag
}

func (c *Controller) NewTag(ctx *gin.Context) {
	var request NewTagRequest

	if err := utils.BindToJson(ctx, &request); err != nil {
		log.Printf("\nError: %+v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	newTag := DomainTag.CreateTag{
		Title: request.Title,
		Color: request.Color,
	}

	domainTagCreated, err := c.ServiceTag.Create(&newTag)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": domainTagCreated,
	})
	return
}

func (c *Controller) GetTagByID(ctx *gin.Context) {

	param, err := strconv.Atoi(ctx.Param("id"))

	if err != nil || param < 1 {

		log.Printf("No es un parámetro válido\n")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No es un parámetro válido",
		})

		return
	}

	tagID := uint(param)

	tag, err := c.ServiceTag.GetById(tagID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": tag,
	})
	return
}

func (c *Controller) GetAllTags(ctx *gin.Context) {

	tag, err := c.ServiceTag.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": tag,
	})
	return
}
