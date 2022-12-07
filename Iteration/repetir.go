package iteration

func Repetir(caractere string, quantidadeRepeticoes int) string{
	var repeticoes string
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}