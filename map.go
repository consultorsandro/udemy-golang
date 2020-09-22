package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[111] = "Maria"
	aprovados[222] = "Pedro"
	aprovados[333] = "João"
	fmt.Println(aprovados[111])

	for cod, nome := range aprovados {
		fmt.Printf("(Cod: %v) %v \n", cod, nome)
	}
	delete(aprovados, 222)
	for cod, nome := range aprovados {
		fmt.Printf("(Cod: %v) %v \n", cod, nome)
	}

	fmt.Println(aprovados)
}
