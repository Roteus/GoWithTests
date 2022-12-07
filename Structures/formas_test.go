package structures

import "testing"

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
	}
}

func TestArea(t *testing.T) {
	t.Run("Área", func(t *testing.T) {
		testesArea := []struct {
			forma    Forma
			temArea float64
		}{
			{forma: Retangulo{Largura: 12, Altura: 6}, temArea: 72.0},
			{forma: Circulo{Raio: 10}, temArea: 314.1592653589793},
			{forma: Triangulo{Base: 12, Altura: 6}, temArea: 36.0},
		}

		for _, tt := range testesArea {
			resultado := tt.forma.Area()
			if resultado != tt.temArea {
				t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, resultado, tt.temArea)
			}
		}
	})
}
