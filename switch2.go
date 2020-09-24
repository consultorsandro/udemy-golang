package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // Procura o primeiro switch que seja verdadeiro
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa Tarde")
	default:
		fmt.Println("Boa noite")
	}

	/* fmt.Println(t)
	hora := time.Now()
	fmt.Println(hora.YearDay())*/
}
