package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/roger/workhub/infraestructure"
	"github.com/roger/workhub/models"
	"github.com/roger/workhub/services"
	"strconv"
)

func CreateAdministratorGroup(c *gin.Context) {
	database := infraestructure.GetConnection()

	var group models.AdministratorGroup

	err := c.BindJSON(&group)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	groupNew, success := services.CreateAdministrador(database, group)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, groupNew)
}

func GetAdministratorGroups(c *gin.Context) {
	database := infraestructure.GetConnection()

	groups, success := services.GetAdministradors(database)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, groups)
}

func GetAdministratorGroupById(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)
	if error != nil {
		c.String(400, "Bad request")
		return
	}

	var group models.AdministratorGroup
	group, success := services.GetAdministradorId(database, idInt)
	if success != true {
		c.String(400, "Error")
		return
	}
	var aux models.AdministratorGroup = models.AdministratorGroup{}
	if group.ID == (aux.ID) {
		c.String(404, "Group not found")
		return
	}

	c.JSON(200, group)
}

func DeleteAdministratorGroup(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)
	if error != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.DeleteAdministrador(database, idInt)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.String(200, message)
}

func UpdateAdministratorGroup(c *gin.Context) {
	database := infraestructure.GetConnection()

	var group models.AdministratorGroup

	err := c.BindJSON(&group)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.UpdateAdministrador(database, group)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.String(200, message)
}
