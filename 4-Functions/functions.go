package main

import "fmt"

func main() {

	soma := somar(10, 20)
	fmt.Println(soma)

	resultSoma, resultSub := calculos(2, 1)

	fmt.Println(resultSoma, resultSub)

	resSoma, _ := calculos(5, 2)
	fmt.Println(resSoma)

	resl := f("teste")

	fmt.Println(resl)

}

func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

var f = func(txt string) string {
	fmt.Println(txt)
	return txt
}
