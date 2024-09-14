package util

import (
	"encoding/json"
	"fmt"
	"mimod/models"
	"os"
)

// guarda un archivo json con las tareas
func TareasAJson(tareas []models.Tarea) error {

	//crea y define el nombre del archivo
	archivo, err := os.Create("tareas_go.json")

	//manejo de errores al crear el archivo
	if err != nil {
		fmt.Println("Error al crear el archivo ", err)
		return err
	}

	defer archivo.Close()

	//prapara el archivo para poder escribir en el.
	encoder := json.NewEncoder(archivo)

	//convierte el slice tareas en formato json y lo escribe en el archivo
	err = encoder.Encode(tareas)

	if err != nil {
		return err
	}

	//si esta todo ok retornamos
	return nil

}
