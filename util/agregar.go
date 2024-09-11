package util

import (
	"bufio"
	"fmt"
	"mimod/models"
	"os"
	"strings"
)

func Agregar(lista []models.Tarea) []models.Tarea {

	//instancia para leer la entrada del usuario
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingrese el titulo de la tarea: ")

	//lee la entrada del usuario
	title, _ := reader.ReadString('\n')

	//borra los espacios de la entrada
	title = strings.TrimSpace(title)

	fmt.Println("Ingrese una descripción: ")

	//lee la entrada del usuario
	description, _ := reader.ReadString('\n')

	//borra los espacios de la entrada
	description = strings.TrimSpace(description) //para borrar espacios antes y despues

	//crea una instancia de tarea
	nuevaTarea := models.Tarea{
		Title:       title,
		Description: description,
		Completed:   false,
	}

	//retorna la lista con la nueva tarea agregada
	return append(lista, nuevaTarea)

}
