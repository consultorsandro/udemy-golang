package main

import "fmt"

func main() {

	i := 1

	var p *int = nil // este ponteiro não tem ainda nenhuma referência

	p = &i // atribui ao ponteiro a referência da variável i

	*p++ /* desreferenciando (usando o * na variável), ou seja, tomando o valor que está na referência
	e incrementando */

	i++ /* incrementando o valor na mesma referência,
	mas altera o valor a partir da ultima alteração, seja no ponteiro ou na variável, por isso o
	resultado é 3. */

	fmt.Println("Se p e i tem a mesma referência:", p == &i)
	fmt.Println("A referência do ponteiro p: ", p)
	fmt.Println("O valor associado à referência *p:", *p)
	fmt.Println("O valor da variável i: ", i)
	fmt.Println("O valor da memória de i: ", &i)

}
