package main

import (
	"errors"
	"fmt"
)

func main() {

	// numeros inteiros
	var n int64 = -10000000
	fmt.Println(n)

	var n2 int = -100000
	fmt.Println(n2)

	n3 := 152554
	fmt.Println(n3)

	//alias = apelidos
	// INT32 = RUNE
	var n4 rune = 123
	fmt.Println(n4)

	//INT8 = byte
	var n5 byte = 123
	fmt.Println(n5)

	// não aceita numero negativo
	var n6 uint = 120
	fmt.Println(n6)

	// numeros reais
	var n7 float32 = 123.5
	fmt.Println(n7)

	var n8 float64 = -1234.45
	fmt.Println(n8)

	n9 := 1545.45
	fmt.Println(n9)

	// strings

	var srt string = "texto"
	fmt.Println(srt)

	// valor da tabela ask

	char := 'a'

	fmt.Println(char)

	// bolleano
	var booleano bool = true
	fmt.Println(booleano)

	//error

	var err error = errors.New("novo error")
	fmt.Println(err)

	//valores zero

	var srt0 string
	fmt.Println(srt0)

	var n0 int
	fmt.Println(n0)

	var float0 float32
	fmt.Println(float0)

	var bool0 bool
	fmt.Println(bool0)

	var erro0 error
	fmt.Println(erro0)

}
