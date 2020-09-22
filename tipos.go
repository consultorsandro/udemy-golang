package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//números inteiros
	fmt.Println(1, 2, 100)
	fmt.Println("Literal inteiro é", reflect.TypeOf(2300))
	/*sem sinal (positivos): uint8 (byte) uint16 (short) uint32 (int) uint64 (long)
	entre parênteses os respectivos em Java */

	var b byte = 255 //alias para o uint8
	fmt.Println("o byte é", reflect.TypeOf(b))

	//com sinal: int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é:", i1)
	fmt.Println("O tipo de i1 é:", reflect.TypeOf(i1))

	var i2 rune = 'a' // posição na tabela unicode
	fmt.Println("Na tabela Unicode:", i2)
	fmt.Println("O rune é", reflect.TypeOf(i2))

	//Valores Reais (float32, float64)
	/* var x = float32(49.99) para converter em float32
	se na declaração não houvesse dedterminado*/
	var x float32 = 49.99

	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // exibe o tipo padrão

	//boolean
	bo := true
	fmt.Println("O tipo de bo é: ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "olá mundo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string s1 é:", len(s1))

	//string com várias linhas, use sinal de crase ( `)
	s2 := `Olá
	meu 
	nome
	é Sandro`

	fmt.Println("O tamanho da String s2 é:", len(s2))

	//caracter único, semelhante a char no Java, use apóstrofo
	char := 'a'
	fmt.Println("A posição de a na tabela Unicode é:", char)
	fmt.Println("o tipo de char é:", reflect.TypeOf(char))

}
