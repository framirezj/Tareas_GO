package auxiliar

import (
	"fmt"
	"mimod/models"
)

func ActualizarEstadoDeTarea(indice int, lista []models.Tarea) []models.Tarea {

	if lista[indice].Completed {
		fmt.Println("Esta tarea ya fue completada.")
	} else {
		lista[indice].Completed = true
		fmt.Println("Listo!, tarea marcada como completada.")
	}

	return lista
}
