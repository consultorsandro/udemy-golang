package main

import "fmt"

func main() {
	// cria o array mas a contagem do índice é feita pelo compilador
	numeros := [...]int{1, 2, 3, 4, 5}

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}
	for _, num := range numeros {
		fmt.Println(num)
	}

}
