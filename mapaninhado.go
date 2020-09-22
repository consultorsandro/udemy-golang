package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"GAbriela": 4000.00,
			"Guga":     3000.00,
		},
		"J": {
			"José": 5000.00,
		},
		"P": {
			"Pedro Junior": 1000.00,
		},
	}
	delete(funcPorLetra, "P")
	for letra, funcionario := range funcPorLetra {
		fmt.Println(letra, funcionario)
	}
}
