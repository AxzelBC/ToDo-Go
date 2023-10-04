package controllerTask

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AxzelBC/ToDo-Go/internal/app/port/IService"
	"github.com/AxzelBC/ToDo-Go/internal/domain/DomainTask"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/utils"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	ServiceTask IService.ServiceTask
}

func (c *Controller) NewTask(ctx *gin.Context) {
	param, err := strconv.Atoi(ctx.Param("listID"))

	if err != nil || param < 1 {

		log.Printf("No es un parámetro válido\n")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No es un parámetro válido",
		})

		return
	}

	listID := uint(param)

	var request NewTaskRequest

	if err := utils.BindToJson(ctx, &request); err != nil {
		log.Printf("\nError: %+v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	newTask := DomainTask.CreateTask{
		Title:       request.Title,
		Description: request.Description,
		StartDate:   request.StartDate,
		DueDate:     request.DueDate,
		Alarm:       request.Alarm,
		Priority:    request.Priority,
		Status:      request.Status,
		ListRefer:   listID,
	}

	domainTaskCreated, err := c.ServiceTask.Create(&newTask)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": domainTaskCreated,
	})
	return
}

func (c *Controller) GetTaskById(ctx *gin.Context) {
	param, err := strconv.Atoi(ctx.Param("id"))

	if err != nil || param < 1 {

		log.Printf("No es un parámetro válido\n")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No es un parámetro válido",
		})

		return
	}

	taskId := uint(param)

	task, err := c.ServiceTask.GetById(taskId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": task,
	})
	return
}

func (c *Controller) GetAllTaks(ctx *gin.Context) {
	param, err := strconv.Atoi(ctx.Param("listID"))

	if err != nil || param < 1 {
		log.Printf("No es un parámetro válido\n")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No es un parámetro válido",
		})

		return
	}

	listID := uint(param)

	tasks, err := c.ServiceTask.GetAll(listID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
	return
}
