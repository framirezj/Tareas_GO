package util

import (
	"fmt"
	"mimod/models"
)

func Listar(lista []models.Tarea) {

	for i, t := range lista {
		fmt.Printf("Indice: %d, Tarea: %s - %s, Completado: %t", i, t.Title, t.Description, t.Completed)

	}

}
