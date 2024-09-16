package auxiliar

import (
	"fmt"
)

func ObtenerEntradaMenu() (int, error) {

	var opcion int

	_, err := fmt.Scanln(&opcion)
	if err != nil {
		return -1, fmt.Errorf("ingrese un número entero: %w", err)
	}

	return opcion, nil

}
