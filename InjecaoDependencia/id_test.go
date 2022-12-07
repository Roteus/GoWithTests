package main

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T) {
	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "globoesportes1")

	resultado := buffer.String()
	esperado := "Olá, globoesportes1"

	if resultado != esperado{
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}