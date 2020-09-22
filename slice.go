package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{1, 2, 3} //slice
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	//Slice é um pedaço de um array. Aponta para uma posição em diante.
	s2 := a2[1:3] // acessa a posição 1 até uma posição antes da posição 3
	fmt.Println(a2, s2)

	s3 := a2[:2] // acessa a posição inicial até uma posição antes da posição 2
	fmt.Println(a2, s3)

	s4 := s2[:1]
	fmt.Println(s2, s4)
}
