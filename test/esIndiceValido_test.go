package test

import (
	"mimod/util/auxiliar"
	"testing"
)

func TestEsIndiceValido(t *testing.T) {

	type Test struct {
		indice     int
		largoLista int
		expected   bool
	}

	casos := []Test{
		{0, 10, true},
		{-1, 10, false},
		{10, 10, false},
		{11, 10, false},
	}

	for _, caso := range casos {
		result := auxiliar.EsIndiceValido(caso.indice, caso.largoLista)
		if result != caso.expected {
			t.Errorf("EsIndiceValido(%d, %d) = %v; expected %v", caso.indice, caso.largoLista, result, caso.expected)
		}
	}

}
