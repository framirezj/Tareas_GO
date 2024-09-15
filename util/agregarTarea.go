package util

import (
	"fmt"
	"mimod/models"
	"mimod/util/auxiliar"
)

func AgregarTarea(lista []models.Tarea) []models.Tarea {

	nuevaTarea, err := auxiliar.CrearTarea()
	if err != nil {
		fmt.Println("Error:", err)
		return lista
	}

	//Retorna la lista con la nueva tarea agregada
	return append(lista, nuevaTarea)

}
