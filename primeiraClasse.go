package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3)) // escrever a variável soma como função e passar os parâmetros

	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(2, 3))
}
