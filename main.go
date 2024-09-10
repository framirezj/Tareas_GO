package main

import (
	"mimod/models"
	"mimod/util"
)

func main() {

	//lista principal
	milista := []models.Tarea{
		{Title: "Levantarse", Description: "5 minutos mas...", Completed: false},
		{Title: "Levantarse 2", Description: "5 minutos mas...", Completed: false},
		{Title: "Levantarse 3", Description: "5 minutos mas...", Completed: false}}

	//crea una nueva tarea para posterior agregarla al listado.
	//miTarea := util.Entradas()

	//agrega nueva tarea a la lista.
	//milista = append(milista, miTarea)

	//lista todas las tareas
	util.Listar(milista)

	//marca como completado una tarea
	milista = util.Completada(milista)

	util.Listar(milista)

}
