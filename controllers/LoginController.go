package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/roger/workhub/infraestructure"
	"github.com/roger/workhub/models"
	"github.com/roger/workhub/services"
)

func Login(c *gin.Context) {
	database := infraestructure.GetConnection()

	var login models.Login

	err := c.BindJSON(&login)
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	loginNew, success := services.Login(database, login)
	if success != true {
		c.String(401, "Error en las credenciales")
		return
	}
	c.JSON(200, loginNew)
}
