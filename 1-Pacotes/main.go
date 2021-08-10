package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá Mundo")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("5245@f.com")
	fmt.Println(erro)
}
