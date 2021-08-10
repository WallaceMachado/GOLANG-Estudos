package auxiliar

import "fmt"

// funções com letras maiusculas são publicas e são visíveis a outros pacotes
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}
