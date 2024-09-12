package main

import (
	"fmt"
	"mimod/models"
	"mimod/util"
)

func main() {

	//lista principal
	tareas := []models.Tarea{}

	//cargar la lista de tareas desde el json que actua como bd
	util.TareasDesdeJson(&tareas)

	//Itera hasta la opción de salida.
	for {

		util.ListarTareas(tareas)
		opcion := util.Menu()

		switch opcion {
		case 1:
			tareas = util.AgregarTarea(tareas)
		case 2:
			util.ListarTareas(tareas)
		case 3:
			tareas = util.CompletarTarea(tareas)
		case 4:
			tareas = util.RemoverTarea(tareas)
		case 5:
			util.EditarTarea(tareas)
		case 6:
			util.TareasAJson(tareas)
		case 7:
			fmt.Println("Hasta pronto!")
			return
		default:
			fmt.Println("Opción inválida")
		}

		util.TareasAJson(tareas)
	}

}
