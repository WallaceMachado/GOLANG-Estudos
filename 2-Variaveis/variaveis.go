package main

import "fmt"

func main() {

	var variavel1 string = "var1"
	fmt.Println(variavel1)

	variavel2 := "var2"

	fmt.Println(variavel2)

	var (
		variavel3 string = "var3"
		variavel4 string = "var4"
	)
	fmt.Println(variavel3, variavel4)

	var6, var7 := "var6", "var7"
	fmt.Println(var6, var7)

	const const1 string = "const1"
	fmt.Println(const1)

	var6, var7 = var7, var6
	fmt.Println(var6, var7)
}
