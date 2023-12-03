package main

import (
	_ "fmt"
	"github.com/roger/workhub/infraestructure"
	"github.com/roger/workhub/models"
	"github.com/roger/workhub/services"
)

func main() {
	infraestructure.Migrate()
	database := infraestructure.GetConnection()

	user := models.User{
		Username: "roger", Password: "123456", Email: "", FirstName: "Roger", LastName: "Gomez", Imagen: nil,
	}
	services.CreateUser(database, user)

}
