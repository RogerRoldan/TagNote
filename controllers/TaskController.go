package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/roger/workhub/infraestructure"
	"github.com/roger/workhub/models"
	"github.com/roger/workhub/services"
)

func CreateTask(c *gin.Context) {
	database := infraestructure.GetConnection()

	var task models.Task

	err := c.BindJSON(&task)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	taskNew, success := services.CreateTask(database, task)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, taskNew)
}

func GetTasks(c *gin.Context) {
	database := infraestructure.GetConnection()

	tasks, success := services.GetTasks(database)

	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, tasks)
}

func GetTaskById(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)
	if error != nil {
		c.String(400, "Bad request")
		return
	}

	var task models.Task
	task, success := services.GetTaskId(database, idInt)
	if success != true {
		c.String(400, "Error")
		return
	}
	var aux models.Task = models.Task{}
	if task.ID == (aux.ID) {
		c.String(404, "Task not found")
		return
	}

	c.JSON(200, task)
}

func DeleteTask(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)
	if error != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.DeleteTask(database, idInt)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.String(200, message)
}

func UpdateTask(c *gin.Context) {
	database := infraestructure.GetConnection()

	var task models.Task

	err := c.BindJSON(&task)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.UpdateTask(database, task)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.String(200, message)
}
