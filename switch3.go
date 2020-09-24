package main

import (
	"fmt"
	"time"
)

// FUNÇÃO que identifica o dado
func tipo(i interface{}) string {
	switch i.(type) { // estrurura que seleciona a resposta conforme dado recebido
	case int:
		return "inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "Nenhum dos anteriores"
	}
}

// Passagem e Impressão dos dados identificados
func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("z"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
