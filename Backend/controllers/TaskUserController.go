package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/roger/workhub/infraestructure"
	"github.com/roger/workhub/models"
	"github.com/roger/workhub/services"
	"strconv"
)

func CreateTaskUser(c *gin.Context) {
	database := infraestructure.GetConnection()

	var taskUser models.TaskUser

	err := c.BindJSON(&taskUser)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	taskUserNew, success := services.CreateTaskUser(database, &taskUser)
	if success != true {
		c.String(400, "Error al asignar tarea a usuario")
		return
	}
	c.JSON(200, taskUserNew)
}

func GetTaskUsers(c *gin.Context) {
	database := infraestructure.GetConnection()

	taskUser, success := services.GetTaskUser(database)

	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, taskUser)
}

func GetTaskUserById(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)

	if error != nil {
		c.String(400, "Bad request")
		return
	}

	var taskUser models.TaskUser
	taskUser, success := services.GetTaskUserId(database, idInt)
	if success != true {
		c.String(400, "Error")
		return
	}

	var aux models.TaskUser = models.TaskUser{}
	if taskUser.ID == aux.ID {
		c.String(400, "Error")
		return
	}

	c.JSON(200, taskUser)
}

func DeleteTaskUser(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)

	if error != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.DeleteTaskUser(database, idInt)

	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, message)
}

func GetTaskUserByUserId(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)
	if error != nil {
		c.String(400, "Bad request")
		return
	}

	taskUsers, success := services.GetTaskUserByUserId(database, idInt)

	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, taskUsers)
}

func GetTaskUserByTaskId(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)
	if error != nil {
		c.String(400, "Bad request")
		return
	}

	taskUsers, success := services.GetTaskUserByTaskId(database, idInt)

	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, taskUsers)
}

func UpdateTaskUser(c *gin.Context) {
	database := infraestructure.GetConnection()

	var taskUser models.TaskUser

	err := c.BindJSON(&taskUser)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.UpdateTaskUser(database, &taskUser)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, message)
}
