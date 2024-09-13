package util

import (
	"bufio"
	"fmt"
	"mimod/models"

	"os"
	"strconv"
	"strings"
)

func CompletarTarea(lista []models.Tarea) []models.Tarea {

	//Instancia para leer la entrada del usuario
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingresa el índice de la tarea completada: ")

	//Lee la entrada del usuario
	entrada, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return lista
	}

	//limpia la entrada de espacios vacíos
	entrada = strings.TrimSpace(entrada)

	//Convierte el string a int
	indice, err := strconv.Atoi(entrada)

	//controlar la entrada del usuario
	if err != nil {
		fmt.Println("Entrada inválida: por favor ingresa un número entero")
		return lista
	}

	//marcar la tarea como completada, si esta dentro del rango.
	if indice >= 0 && indice < len(lista) {

		if lista[indice].Completed {
			fmt.Println("Esta tarea ya fue completada.")
		} else {
			lista[indice].Completed = true
			fmt.Println("Listo!, tarea marcada como completada.")
		}

	} else {
		fmt.Println("Índice fuera de rango.")
		return lista
	}

	return lista

}
