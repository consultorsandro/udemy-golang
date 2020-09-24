package main

import "fmt"

// Trocar define a ordem que está no retorno
func trocar(p1, p2 int) (seg int, pri int) {
	seg = p2 // o primeiro retorno vai receber o segundo parâmetro
	pri = p1 // o segundo retorno vai receber o primeiro

	return

}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	seg, pri := trocar(2, 3)
	fmt.Println(seg, pri)

	a, b := trocar(3, 2)
	fmt.Println(a, b)

}
