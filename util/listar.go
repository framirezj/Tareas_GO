package util

import (
	"fmt"
	"mimod/models"
)

func Listar(lista []models.Tarea) {

	fmt.Println("\n###############")
	fmt.Println("LISTA DE TAREAS")
	fmt.Println("---------------")

	for i, tarea := range lista {
		fmt.Printf("Indice: %d, Tarea: %s - %s, Completado: %t\n", i, tarea.Title, tarea.Description, tarea.Completed)

	}
	fmt.Println("---------------")
	fmt.Println("###############\n")

}
