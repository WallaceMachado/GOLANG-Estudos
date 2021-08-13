ackage main

import (
	"fmt"
	"time"
)

func main() {
	// executar mas não esperar a conclusão para seguir com o programa
	go escrever("Olá Mundo!") // goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}