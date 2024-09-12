package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Menu() int {

	//instancia para leer entrada de usuario.
	reader := bufio.NewReader(os.Stdin)

	//lanza el men[u] de opciones
	fmt.Println(`
Menú:
1. Agregar Tarea
2. Listar Tareas
3. Marcar Tarea Completada
4. Remover Tarea
5. Actualizar Tarea
6. Guarda las tareas en un archivo Json
7. Salir
Seleccione una opción:`)

	//leer la entrada de usuario
	entrada, _ := reader.ReadString('\n')

	//borra los espacios antes y después de la entrada
	entrada = strings.TrimSpace(entrada)

	//convierte la entrada string a int
	opcion, _ := strconv.Atoi(entrada)

	//devuelve la opci[o]n escogida
	return opcion
}
