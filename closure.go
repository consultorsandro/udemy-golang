package main

import "fmt"

func closure() func() { // a função retorna uma função
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}
func main() {
	x := 20
	fmt.Println(x)

	exibeX := closure()
	exibeX()
}
