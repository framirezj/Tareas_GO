package util

import (
	"bufio"
	"fmt"
	"mimod/models"
	"os"
	"strings"
)

func Entradas() models.Tarea {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingrese el titulo de la tarea: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title) //para borrar espacios antes y despues

	fmt.Println("Ingrese una descripción: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description) //para borrar espacios antes y despues

	return models.Tarea{
		Title:       title,
		Description: description,
		Completed:   false,
	}

}
