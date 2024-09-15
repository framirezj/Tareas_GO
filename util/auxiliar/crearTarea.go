package auxiliar

import (
	"fmt"
	"mimod/models"
)

func CrearTarea() (models.Tarea, error) {

	fmt.Println("Ingrese el título de la tarea:")
	title, err := ObtenerEntrada()
	if err != nil {
		fmt.Println("Error:", err)
		return models.Tarea{}, err
	}

	fmt.Println("Ingrese una descripción:")
	description, err := ObtenerEntrada()
	if err != nil {
		fmt.Println("Error:", err)
		return models.Tarea{}, err
	}

	//Crea una instancia de tarea
	return models.Tarea{
		Title:       title,
		Description: description,
		Completed:   false,
	}, nil

}
