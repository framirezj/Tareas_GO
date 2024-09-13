package test

import (
	"io"
	"mimod/util/auxiliar"
	"os"
	"strings"
	"testing"
)

/*
en esta prueba tendre que crear un archivo de texto
temporal que simulara la entrada del usuario.
*/

func TestObtenerIndiceDesdeEntrada(t *testing.T) {

	//respalda el valor original de os.stdin
	originalStdin := os.Stdin

	//crear el archivo temporal
	tempFile, err := os.CreateTemp("", "test_entrada")
	if err != nil {
		t.Fatalf("no se pudo crear un archivo temporal: %v", err)
	}

	//borrar el archivo al salir del test
	defer os.Remove(tempFile.Name())

	//redirige el archivo temporal a la entrada
	os.Stdin = tempFile

	//-----------------------------------------------

	//Caso 1: Entrada Válida

	t.Run("Entrada válida", func(t *testing.T) {

		//simular la entrada en el archivo temporal
		tempFile.WriteString("42\n")
		tempFile.Seek(0, io.SeekStart)

		//ejecuta la funcion
		indice, err := auxiliar.ObtenerIndiceDesdeEntrada()

		//varificar el resultado
		if err != nil {
			t.Fatalf("se produjo un error inesperado: %v", err)
		}

		if indice != 42 {
			t.Fatalf("Se esperaba 42, pero se obtuvo %d", indice)
		}

		//volver el archivo al estado inicial
		tempFile.Truncate(0)
	})

	//Caso 2:
	t.Run("Entrada no numérica", func(t *testing.T) {

		//simular la entrada en el archivo temporal
		tempFile.WriteString("no_numero\n")
		tempFile.Seek(0, io.SeekStart)

		//ejecutar la función
		_, err := auxiliar.ObtenerIndiceDesdeEntrada()

		if err == nil {
			t.Fatalf("se esperaba un error, pero no se produjo ningún error")
		}

		expectedErr := "entrada inválida: por favor ingresa un número entero"

		if !strings.Contains(err.Error(), expectedErr) {
			t.Errorf(
				"se esperaba el error '%s', pero se obtuvo '%s'", expectedErr, err.Error())
		}

		//volver el archivo al estado inicial
		tempFile.Truncate(0)

	})

	//restaurar el os.Stdin
	os.Stdin = originalStdin

}
