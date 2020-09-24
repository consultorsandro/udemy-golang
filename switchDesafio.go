package main

import "fmt"

func conceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "Conceito A"
	case n >= 7 && n <= 8.9:
		return "Conceito B"
	case n >= 5 && n <= 6.9:
		return "Conceito C"
	case n >= 3 && n <= 4.9:
		return "Conceito D"
	default:
		return "Conceito E"
	}
}

func main() {
	fmt.Println(conceito(9.8))
	fmt.Println(conceito(7.0))
	fmt.Println(conceito(2.8))
}
