package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const esperanto = "esperanto"
const prefixoOlaEsperanto = "Saluton, "
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}
	return prefixoSaudacao(idioma) + nome
}

func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case esperanto:
		prefixo = prefixoOlaEsperanto
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func main() {
	fmt.Println(Ola("mundo", prefixoOlaPortugues))
}
