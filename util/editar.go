package util

import (
	"bufio"
	"fmt"
	"mimod/models"
	"os"
	"strconv"
	"strings"
)

func Editar(lista []models.Tarea) []models.Tarea {

	//instancia del lector
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingresa el indice de la tarea que deseas actualizar: ")

	//lee la entrada del usuario
	entrada, _ := reader.ReadString('\n')

	//limpia la entrada de espacios vacios
	entrada = strings.TrimSpace(entrada)

	//convierte el string a int
	indice, _ := strconv.Atoi(entrada)

	fmt.Println("Ingrese el titulo de la tarea: ")

	//lee la entrada del usuario
	title, _ := reader.ReadString('\n')

	//borra los espacios de la entrada
	title = strings.TrimSpace(title)

	fmt.Println("Ingrese una descripción: ")

	//lee la entrada del usuario
	description, _ := reader.ReadString('\n')

	//borra los espacios de la entrada
	description = strings.TrimSpace(description)

	//mostrar la tarea
	tarea := models.Tarea{
		Title:       title,
		Description: description,
		Completed:   true,
	}

	lista[indice] = tarea

	return lista

}
