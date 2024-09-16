package util

import (
	"fmt"
	"mimod/models"
	"mimod/util/auxiliar"
)

func ActualizarTarea(lista []models.Tarea) []models.Tarea {

	indice, err := auxiliar.ObtenerIndiceDesdeEntrada()
	if err != nil {
		fmt.Println("Error:", err)
		return lista
	}

	tarea, err := auxiliar.CrearTarea()
	if err != nil {
		fmt.Println("Error:", err)
		return lista
	}

	lista[indice] = tarea

	return lista

}
