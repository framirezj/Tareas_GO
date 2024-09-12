package util

import (
	"bufio"
	"fmt"
	"mimod/models"
	"os"
	"strings"
)

func AgregarTarea(lista []models.Tarea) []models.Tarea {

	//Instancia para leer la entrada del usuario
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingrese el título de la tarea: ")

	//Lee la entrada del usuario
	title, _ := reader.ReadString('\n')

	//Quita los espacios vacíos de la entrada
	title = strings.TrimSpace(title)

	fmt.Println("Ingrese una descripción: ")

	//Lee la entrada del usuario
	description, _ := reader.ReadString('\n')

	//Quita los espacios de la entrada
	description = strings.TrimSpace(description)

	//Crea una instancia de tarea
	nuevaTarea := models.Tarea{
		Title:       title,
		Description: description,
		Completed:   false,
	}

	//Retorna la lista con la nueva tarea agregada
	return append(lista, nuevaTarea)

}
