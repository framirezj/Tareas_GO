package util

import (
	"fmt"
	"mimod/models"
)

func Listar(lista []models.Tarea) {

	fmt.Println("\n###############")
	fmt.Println("LISTA DE TAREAS")
	fmt.Println("---------------")

	//Iteración del slice de tareas
	for i, tarea := range lista {
		fmt.Printf("Indice: %d, \tTarea: %s, \tDescripción: %s, \tCompletado: %t\n", i, tarea.Title, tarea.Description, tarea.Completed)
	}

	fmt.Println("---------------")
	fmt.Println("################")

}
