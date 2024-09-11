package util

import (
	"bufio"
	"fmt"
	"mimod/models"
	"os"
	"strconv"
	"strings"
)

func Remover(lista []models.Tarea) []models.Tarea {

	//instancia del lector
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingresa el índice de la tarea que deseas remover: ")

	//lee la entrada del usuario
	entrada, _ := reader.ReadString('\n')

	//limpia la entrada de espacios vacíos
	entrada = strings.TrimSpace(entrada)

	//convierte el string a int
	indice, _ := strconv.Atoi(entrada)

	//crea una nueva lista, pero sin el elemento con la posición del índice.
	//[:indice] considera el primer índice y uno antes del indicado
	//[indice+1:] considera del siguiente indice indicado hasta el último elemento de la lista
	return append(lista[:indice], lista[indice+1:]...)

}
