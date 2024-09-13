package auxiliar

//verifica si el índice esta dentro del rango válido de la lista.
func EsIndiceValido(indice, largoLista int) bool {

	return indice >= 0 && indice < largoLista

}
