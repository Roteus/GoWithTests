package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const escrita = "escrita"
const pausa = "pausa"

type SpyContagemOperacoes struct{
	Chamadas []string
}

type TempoSpy struct{
	duracaoPausa time.Duration
}

func (t *TempoSpy) Pausa(duracao time.Duration){
	t.duracaoPausa = duracao
}

func (s *SpyContagemOperacoes) Pausa(){
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error){
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

func TestContagem(t *testing.T) {

t.Run("imprime 3 até Go!", func(t *testing.T) {
	buffer := &bytes.Buffer{}
	Contagem(buffer, &SpyContagemOperacoes{})

	resultado := buffer.String()
	esperado := "3\n2\n1\nGo!"

	if resultado != esperado{
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}

})

	t.Run("pausa antes de cada impresssão", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		resultado := buffer.String()
		spyImpressoraSleep := &SpyContagemOperacoes{}
		Contagem(spyImpressoraSleep, spyImpressoraSleep)

		esperado := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(esperado, spyImpressoraSleep.Chamadas){
			t.Errorf("esperado %v chamadas, resultado %v", esperado, resultado)
		}
	})
}

func TestSleeperConfiguravel(t *testing.T){
	tempoPausa := 5 * time.Second

	tempoSpy := &TempoSpy{}
	sleeper := SleeperConfiguravel{tempoPausa, tempoSpy.Pausa}
	sleeper.Pausa()

	if tempoSpy.duracaoPausa != tempoPausa{
		t.Errorf("deveria ter passado %v, mas pausou por %v", tempoPausa, tempoSpy)
	}
}