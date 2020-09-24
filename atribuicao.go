package main

import (
	"fmt"
)

func main() {

	var z byte = 30
	fmt.Println(z)

	//Atribuições curtas e especiais
	a := 5 // declara, atribui e autodetermina o tipo
	fmt.Println("Usando :=", a)
	b := 5
	b += 5 // b = b + 5
	fmt.Println("Usando +=", b)
	c := 5
	c -= 5 // c = c - 5
	fmt.Println("Usando -=", c)
	d := 5
	d /= 5 // d = d / 5
	fmt.Println("Usando /=", d)
	e := 5
	e *= 5 // e = e * 5
	fmt.Println("Usando *=", e)
	f := 5
	f %= 5 // f = f % 5
	fmt.Println("Usando %=", f)

	// Troca rápida de valores entre duas variáveis
	x, y := 1, 2
	fmt.Println("Valores de x e y:", x, y)
	x, y = y, x
	fmt.Println("Depois da troca:", x, y)

}
