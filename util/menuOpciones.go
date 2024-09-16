package util

import (
	"fmt"
	"mimod/util/auxiliar"
)

func Menu() int {

	//lanza el men[u] de opciones
	fmt.Println(`
Menú:
1. Agregar Tarea
2. Listar Tareas
3. Marcar Tarea como completada
4. Remover Tarea
5. Actualizar Tarea
6. Salir
Seleccione una opción:`)

	opcion, err := auxiliar.ObtenerEntradaMenu()
	if err != nil {
		fmt.Println("Error:", err)
		Menu()
	}

	//devuelve la opci[o]n escogida
	return opcion
}
