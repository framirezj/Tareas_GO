package util

import (
	"bufio"
	"fmt"
	"mimod/models"

	"os"
	"strconv"
	"strings"
)

func Completada(lista []models.Tarea) []models.Tarea {

	//instancia del lector
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingresa el indice de la tarea completada: ")

	//lee la entrada del usuario
	entrada, _ := reader.ReadString('\n')

	//limpia la entrada de espacios vacios
	entrada = strings.TrimSpace(entrada)

	//convierte el string a int
	indice, _ := strconv.Atoi(entrada)

	//marca la tarea como completada
	lista[indice].Completed = true

	return lista

}
