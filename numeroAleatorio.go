package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(26)
}

func main() {
	i := numeroAleatorio()
	fmt.Println("O número aleatório é:", i)

	if i > 10 { // estrutura com dupla condição
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu")
	}
}
