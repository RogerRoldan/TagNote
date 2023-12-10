package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/roger/workhub/infraestructure"
	"github.com/roger/workhub/models"
	"github.com/roger/workhub/services"
	"strconv"
)

func CreateGroupUser(c *gin.Context) {
	database := infraestructure.GetConnection()

	var groupUser models.GroupUser

	err := c.BindJSON(&groupUser)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	groupUserNew, success := services.CreateGroupUser(database, groupUser)
	if success != true {
		c.String(400, "Error al asignar grupo al usuario")
		return
	}
	c.JSON(200, groupUserNew)
}

func GetGroupUsers(c *gin.Context) {
	database := infraestructure.GetConnection()

	groupUsers, success := services.GetGroupUsers(database)

	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, groupUsers)
}

func GetGroupUserById(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)

	if error != nil {
		c.String(400, "Bad request")
		return
	}

	var groupUser models.GroupUser
	groupUser, success := services.GetGroupUserId(database, idInt)
	if success != true {
		c.String(400, "Error")
		return
	}

	var aux models.GroupUser = models.GroupUser{}
	if groupUser.ID == (aux.ID) {
		c.String(400, "Error")
		return
	}

	c.JSON(200, groupUser)
}

func DeleteGroupUser(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)

	if error != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.DeleteGroupUser(database, idInt)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.String(200, message)
}

func GetGroupUsersByUserId(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)
	if error != nil {
		c.String(400, "Bad request")
		return
	}

	groupUsers, success := services.GetGroupUsersByUserId(database, idInt)

	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, groupUsers)
}

func GetGroupUsersByGroupId(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)
	if error != nil {
		c.String(400, "Bad request")
		return
	}

	groupUsers, success := services.GetGroupUsersByGroupId(database, idInt)

	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, groupUsers)
}

func UpdateGroupUser(c *gin.Context) {
	database := infraestructure.GetConnection()

	var groupUser models.GroupUser

	err := c.BindJSON(&groupUser)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.UpdateGroupUser(database, groupUser)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.String(200, message)
}
