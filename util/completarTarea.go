package util

import (
	"fmt"
	"mimod/models"
	"mimod/util/auxiliar"
)

func CompletarTarea(lista []models.Tarea) []models.Tarea {

	//solicitar el indice de la tarea.
	indice, err := auxiliar.ObtenerIndiceDesdeEntrada()
	if err != nil {
		fmt.Println("Error:", err)
		return lista
	}

	//marcar la tarea como completada, si esta dentro del rango.
	if auxiliar.EsIndiceValido(indice, len(lista)) {
		//actualiza la lista
		lista = auxiliar.ActualizarEstadoDeTarea(indice, lista)
	} else {
		fmt.Println("Índice fuera de rango válido.")
	}

	return lista

}
