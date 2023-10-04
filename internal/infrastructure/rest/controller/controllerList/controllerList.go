package controllerList

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AxzelBC/ToDo-Go/internal/app/port/IService"
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainList"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/utils"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	ServiceList IService.ServiceList
}

func (c *Controller) NewList(ctx *gin.Context) {
	var request NewListRequest

	if err := utils.BindToJson(ctx, &request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	newList := DomainList.CreateList{
		Title:  request.Title,
		Status: request.Status,
		Tasks:  request.Task,
	}

	domainListCreated, err := c.ServiceList.Create(&newList)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": domainListCreated,
	})
	return
}

func (c *Controller) GetListByID(ctx *gin.Context) {

	param, err := strconv.Atoi(ctx.Param("id"))

	if err != nil || param < 1 {

		log.Printf("\nNo es un parámetro válido\n")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No es un parámetro válido",
		})

		return
	}

	idList := uint(param)

	list, err := c.ServiceList.GetById(idList)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": list,
	})
	return
}

func (c *Controller) GetAllLists(ctx *gin.Context) {

	list, err := c.ServiceList.GetAll()

	if err != nil {

		log.Printf("\n%+v\n", err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": list,
	})
	return
}
