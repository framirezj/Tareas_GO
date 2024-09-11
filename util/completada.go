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

	//Instancia para leer la entrada del usuario
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingresa el índice de la tarea completada: ")

	//Lee la entrada del usuario
	entrada, _ := reader.ReadString('\n')

	//limpia la entrada de espacios vacíos
	entrada = strings.TrimSpace(entrada)

	//Convierte el string a int
	indice, _ := strconv.Atoi(entrada)

	//Marca la tarea como completada
	lista[indice].Completed = true

	return lista

}
