package main

import "fmt"

func main() {
	s := make([]int, 10) // criar o array a partir do slice, determinando inclusive o tamanho do array igual ao slice
	s[9] = 12
	fmt.Println(s)

	// criar o array a partir do slice, determinando inclusive a capacidade do array maior que o slice(20)
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s[19])

	s = append(s, 89)
	fmt.Println(s, len(s), cap(s))
}
