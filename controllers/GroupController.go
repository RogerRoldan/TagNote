package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/roger/workhub/infraestructure"
	"github.com/roger/workhub/models"
	"github.com/roger/workhub/services"
	"strconv"
)

func CreateGroup(c *gin.Context) {
	database := infraestructure.GetConnection()

	var group models.Group

	err := c.BindJSON(&group)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	groupNew, success := services.CreateGroup(database, group)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, groupNew)
}

func GetGroups(c *gin.Context) {
	database := infraestructure.GetConnection()

	groups, success := services.GetGroups(database)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, groups)
}

func GetGroupById(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)
	if error != nil {
		c.String(400, "Bad request")
		return
	}

	var group models.Group
	group, success := services.GetGroupId(database, idInt)
	if success != true {
		c.String(400, "Error")
		return
	}
	var aux models.Group = models.Group{}
	if group.ID == (aux.ID) {
		c.String(404, "Group not found")
		return
	}

	c.JSON(200, group)
}

func DeleteGroup(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)
	if error != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.DeleteGroup(database, idInt)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.String(200, message)
}

func UpdateGroup(c *gin.Context) {
	database := infraestructure.GetConnection()

	var group models.Group

	err := c.BindJSON(&group)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.UpdateGroup(database, group)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.String(200, message)
}
