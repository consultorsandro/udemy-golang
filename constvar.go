package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // o copilador já inferiu o tipo

	//forma reduzida de declarar e atribuir valor a uma variável
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	//declaração comum
	var e, f bool = true, false
	fmt.Println(e, f)

	//declaração curta
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)

}
