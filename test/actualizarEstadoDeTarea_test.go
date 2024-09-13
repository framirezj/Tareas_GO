package test

import (
	"mimod/models"
	"mimod/util/auxiliar"
	"testing"
)

func TestActualizarEstadoDeTarea(t *testing.T) {

	//struct
	type Test struct {
		name          string
		indiceEntrada int
		listaEntrada  []models.Tarea
		expectedLista []models.Tarea
	}

	//casos

	tests := []Test{
		{
			name:          "Caso 1: Completar tarea no completada",
			indiceEntrada: 0,
			listaEntrada: []models.Tarea{
				{Title: "Reunión", Description: "14:00 hrs", Completed: false},
				{Title: "Reunión", Description: "20:00 hrs", Completed: false},
			},

			expectedLista: []models.Tarea{
				{Title: "Reunión", Description: "14:00 hrs", Completed: true},
				{Title: "Reunión", Description: "20:00 hrs", Completed: false},
			},
		},
		{
			name:          "Caso 2: No completar tarea ya completada",
			indiceEntrada: 1,
			listaEntrada: []models.Tarea{
				{Title: "Reunión", Description: "14:00 hrs", Completed: true},
				{Title: "Reunión", Description: "20:00 hrs", Completed: false},
			},

			expectedLista: []models.Tarea{
				{Title: "Reunión", Description: "14:00 hrs", Completed: true},
				{Title: "Reunión", Description: "20:00 hrs", Completed: true},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			//ejecutar la funcion
			result := auxiliar.ActualizarEstadoDeTarea(test.indiceEntrada, test.listaEntrada)

			if result[0] != test.expectedLista[0] {
				t.Errorf("ActualizarEstadoDeTarea() = %v, esperado %v", result, test.expectedLista)
			}

		})
	}

}
