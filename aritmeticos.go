package main

import (
	"fmt"
	"math"
)

func main() {

	a := 3
	b := 2
	// Operações usando valores inteiros
	fmt.Println("Soma ->", a+b)
	fmt.Println("Subtração->", a-b)
	fmt.Println("Divisão ->", a/b)
	fmt.Println("Multiplicação ->", a*b)
	fmt.Println("Módulo ou resto->", a%b)

	//operações com métodos geralmente usam float64
	c := 3.0
	d := 2.0

	fmt.Println("Maior -> ", math.Max(float64(a), float64(b)))
	fmt.Println("Menor -> ", math.Min(float64(a), float64(b)))
	fmt.Println("Operação de Potência -> ", math.Pow(c, d))

}
