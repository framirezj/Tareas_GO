package util

import (
	"encoding/json"
	"fmt"
	"mimod/models"
	"os"
	"time"
)

// guarda un archivo json con las tareas
func TareasAJson(tareas []models.Tarea) error {

	//crea y define el nombre del archivo
	//el formato del archivo sera ejemplo: archivo-fechahora.json

	fechaHora := time.Now()
	formatoFechaHora := fechaHora.Format("2006-01-02_15-04-05")

	nombreArchivo := fmt.Sprintf("archivo-%s.json", formatoFechaHora)

	//crea el archivo vacio con el nombre previamente definido
	archivo, err := os.Create(nombreArchivo)

	//manejo de errores al crear el archivo
	if err != nil {
		fmt.Println("Error al crear el archivo ", err)
		return err
	}

	//prapara el archivo para poder escribir en el.
	encoder := json.NewEncoder(archivo)

	//convierte el slice tareas en formato json y lo escribe en el archivo
	err = encoder.Encode(tareas)

	if err != nil {
		return err
	}

	//si esta todo ok cerramos y retornamos
	archivo.Close()
	return nil

}
