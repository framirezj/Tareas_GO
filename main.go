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

	util.ListarTareas(tareas)

	//Itera hasta la opción de salida.
	for {

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
			fmt.Println("Hasta pronto!")
			return
		default:
			fmt.Println("Opción inválida")
			return
		}

		util.TareasAJson(tareas)
	}

}
