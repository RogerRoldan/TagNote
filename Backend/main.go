package main

import (
	_ "fmt"
	"github.com/roger/workhub/infraestructure"
	"github.com/roger/workhub/routes"
)

func main() {
	infraestructure.Migrate()

	routes.Init()

}
