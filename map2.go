package main

import "fmt"

func main() {

	funcEmpregados := map[string]float64{
		"José":  2000.45,
		"Gabi":  3000.10,
		"Pedro": 5000.01,
	}
	funcEmpregados["Rafael"] = 1500.02
	fmt.Println(funcEmpregados)

	for nome, salario := range funcEmpregados {
		fmt.Println(nome, salario)
	}
}
