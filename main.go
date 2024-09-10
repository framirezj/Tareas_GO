package main

import (
	"mimod/models"
	"mimod/util"
)

func main() {

	milista := []models.Tarea{}

	miTarea := util.Entradas()

	milista = append(milista, miTarea)

	util.Listar(milista)

}
