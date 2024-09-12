package util

import (
	"encoding/json"
	"mimod/models"
	"os"
)

func TareasDesdeJson(tareas *[]models.Tarea) error {

	//buscar el archivo
	archivo, err := os.Open("tareas_go.json")

	if err != nil {
		return err
	}

	decoder := json.NewDecoder(archivo)

	err = decoder.Decode(tareas)

	if err != nil {
		return err
	}

	archivo.Close()
	return nil

}
