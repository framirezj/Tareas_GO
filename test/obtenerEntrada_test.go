package test

import (
	"io"
	"mimod/util/auxiliar"
	"os"
	"testing"
)

func TestObtenerEntrada(t *testing.T) {

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

	//Caso 1: Entrada válida
	t.Run("Entrada válida", func(t *testing.T) {

		//simular la entrada en el archivo temporal
		tempFile.WriteString("hola\n")
		tempFile.Seek(0, io.SeekStart)

		//ejecuta la función
		result, err := auxiliar.ObtenerEntrada()

		if err != nil {
			t.Fatalf("se produjo un error inesperado: %v", err)
		}

		if result != "hola" {
			t.Fatalf("se esperaba 'hola', pero se obtuvo %s", result)
		}

		//volver el archivo al estado inicial
		tempFile.Truncate(0)
		tempFile.Seek(0, io.SeekStart)

	})

	//Caso 2: Entrada vacía
	t.Run("Entrada vacía", func(t *testing.T) {

		//simular la entrada en el archivo temporal
		string := "\n"
		tempFile.WriteString(string)
		tempFile.Seek(0, io.SeekStart)

		//ejecuta la función
		_, err := auxiliar.ObtenerEntrada()

		expectedMessage := "la entrada no puede estar vacía"

		if err.Error() != expectedMessage {
			t.Errorf("esperaba el error '%v', pero obtuvo: %v", expectedMessage, err.Error())
		}

		//volver el archivo al estado inicial
		tempFile.Truncate(0)

	})

	os.Stdin = originalStdin

}
