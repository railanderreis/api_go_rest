package main

import (
	"fmt"

	"github.com/railanderreis/go-rest-api/models"
	"github.com/railanderreis/go-rest-api/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{ID: 1, Nome: "nome 1", Historia: "historia 1"},
		{ID: 2, Nome: "nome 2", Historia: "historia 2"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
