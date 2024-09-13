package auxiliar

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ObtenerIndiceDesdeEntrada() (int, error) {

	//Instancia para leer la entrada del usuario
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingresa el índice de la tarea completada: ")

	//Lee la entrada del usuario
	entrada, err := reader.ReadString('\n')
	if err != nil {
		return -1, fmt.Errorf("error al leer la entrada: %w", err)
	}

	//limpia la entrada de espacios vacíos
	entrada = strings.TrimSpace(entrada)

	//Convierte el string a int
	indice, err := strconv.Atoi(entrada)

	//controlar la entrada del usuario
	if err != nil {
		return -1, fmt.Errorf("entrada inválida: por favor ingresa un número entero")
	}

	return indice, nil

}
