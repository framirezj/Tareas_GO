package main

import (
	"fmt"
	"mimod/models"
	"mimod/util"
)

func main() {

	//lista principal
	tareas := []models.Tarea{
		{Title: "Levantarse", Description: "5 minutos más...", Completed: false},
		{Title: "Levantarse 2", Description: "5 minutos más...", Completed: false},
		{Title: "Levantarse 3", Description: "5 minutos más...", Completed: false}}

	//Itera hasta la opción de salida.
	for {
		util.Listar(tareas)
		switch util.Menu() {
		case 1:
			tareas = util.Agregar(tareas)
		case 2:
			util.Listar(tareas)
		case 3:
			tareas = util.Completada(tareas)
		case 4:
			tareas = util.Remover(tareas)
		case 5:
			util.Editar(tareas)
		case 6:
			fmt.Println("Hasta pronto!")
			return
		default:
			fmt.Println("Opción inválida")
		}
	}

}
