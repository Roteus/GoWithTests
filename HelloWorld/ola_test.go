package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris", "")
		esperado := "Olá, Chris"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T){
		resultado := Ola("Chris", "espanhol")
		esperado := "Hola, Chris"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T){
		resultado := Ola("Chris", "frances")
		esperado := "Bonjour, Chris"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em esperanto", func(t *testing.T){
		resultado := Ola("Chris", "esperanto")
		esperado := "Saluton, Chris"

		verificaMensagemCorreta(t, resultado, esperado)
	})
}