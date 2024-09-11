package main

import (
	"fmt"
	"mimod/models"
	"mimod/util"
)

func main() {

	//lista principal
	milista := []models.Tarea{
		{Title: "Levantarse", Description: "5 minutos mas...", Completed: false},
		{Title: "Levantarse 2", Description: "5 minutos mas...", Completed: false},
		{Title: "Levantarse 3", Description: "5 minutos mas...", Completed: false}}

	//iteración hasta la opción de salida.
	for {
		switch util.Menu() {
		case 1:
			milista = util.Agregar(milista)
		case 2:
			util.Listar(milista)
		case 3:
			milista = util.Completada(milista)
		case 4:
			milista = util.Remover(milista)
		case 5:
			fmt.Println("Hasta pronto!")
			return
		default:
			fmt.Println("Opción invalida")
		}
	}

}
