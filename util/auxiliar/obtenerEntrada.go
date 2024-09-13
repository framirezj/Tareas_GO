package auxiliar

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ObtenerEntrada() (string, error) {

	//Instancia para leer la entrada del usuario
	reader := bufio.NewReader(os.Stdin)

	//Lee la entrada del usuario
	entrada, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error al leer la entrada: %w", err)
	}

	//Quita los espacios vacíos de la entrada
	entrada = strings.TrimSpace(entrada)

	//prevenir entradas vacias
	if entrada == "" {
		return "", fmt.Errorf("la entrada no puede estar vacía")
	}

	return entrada, nil

}
