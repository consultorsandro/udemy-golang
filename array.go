package main

import "fmt"

func main() {
	/*var notas [3]float64 // array é uma estrutura homogênea(tipo) e estática(não expansível)
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1

	fmt.Println(notas)*/
	notas := [3]float64{7.8, 4.3, 9.1}

	// calcular a média das notas após percorrer o array
	total := 0.0
	for i := 0; i < len(notas); i++ { // len dá o limite da condição, a quantidade de posições do array(3)
		// recebe a soma das notas a cada iteração
		total += notas[i]
	} // len exibe a quantidade de posições
	media := total / float64(len(notas)) // converte o denominador para float64 por causa do tipo do array
	fmt.Printf("Média %.2f\n", media)
}
