package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

//getNomeCompleto une os nomes da estrutc pessoa na ordem configurada
func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// setNomeComleto altera os dados da struct pessoa na variável onde está referenciada.
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]

}
func main() {
	p1 := pessoa{"Pedro", "Silva"}    // struct pessoa recebe os dados
	fmt.Println(p1.getNomeCompleto()) // exibe os nomes na ordem configura

	p1.setNomeCompleto("Maria Pereira") // altera os dados da struct pessoa na ordem configurada
	fmt.Println(p1.getNomeCompleto())   // exibe os novos nomes depois de alterados, na ordem configurada
}
