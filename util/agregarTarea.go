package util

import (
	"fmt"
	"mimod/models"
	"mimod/util/auxiliar"
)

func AgregarTarea(lista []models.Tarea) []models.Tarea {

	fmt.Println("Ingrese el título de la tarea:")
	title, err := auxiliar.ObtenerEntrada()
	if err != nil {
		fmt.Println("Error:", err)
		return lista
	}

	fmt.Println("Ingrese una descripción:")
	description, err := auxiliar.ObtenerEntrada()
	if err != nil {
		fmt.Println("Error:", err)
		return lista
	}

	//Crea una instancia de tarea
	nuevaTarea := models.Tarea{
		Title:       title,
		Description: description,
		Completed:   false,
	}

	//Retorna la lista con la nueva tarea agregada
	return append(lista, nuevaTarea)

}
