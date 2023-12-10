package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/roger/workhub/infraestructure"
	"github.com/roger/workhub/models"
	"github.com/roger/workhub/services"
	"strconv"
)

func CreateUser(c *gin.Context) {
	database := infraestructure.GetConnection()

	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	userNew, success := services.CreateUser(database, user)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, userNew)
}

func GetUsers(c *gin.Context) {
	database := infraestructure.GetConnection()

	users, success := services.GetUsers(database)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.JSON(200, users)

}

func GetUserById(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)
	if error != nil {
		c.String(400, "Bad request")
		return
	}

	var user models.User
	user, success := services.GetUserId(database, idInt)
	if success != true {
		c.String(400, "Error")
		return
	}
	var aux models.User = models.User{}
	if user.ID == (aux.ID) {
		c.String(404, "User not found")
		return
	}

	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	database := infraestructure.GetConnection()

	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.UpdateUser(database, user)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.String(200, message)
}

func DeleteUser(c *gin.Context) {
	database := infraestructure.GetConnection()

	id := c.Param("id")
	idInt, error := strconv.Atoi(id)
	if error != nil {
		c.String(400, "Bad request")
		return
	}

	message, success := services.DeleteUser(database, idInt)
	if success != true {
		c.String(400, "Error")
		return
	}
	c.String(200, message)
}
