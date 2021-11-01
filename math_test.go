package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: %d. \nEsperado: %d\n", total, 30)
	}
}
