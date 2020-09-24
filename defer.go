package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim")
	if numero > 5000 {
		fmt.Println("Valor Alto...")
		return 5000
		//If dispensa o bloco 'else' final, como usado na linguagem C
	} /* else {
	fmt.Println("Valor baixo.")
	return numero
	}*/
	fmt.Println("Valor baixo")
	return numero

}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(1000))
}
