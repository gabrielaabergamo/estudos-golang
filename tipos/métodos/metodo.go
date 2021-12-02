package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Gon", "Freecss"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Killua Zoldyck")
	fmt.Println(p1.getNomeCompleto())
	//no método setNomeCompleto o receptor é um ponteiro do tipo pessoa, logo estamos
	//alterando os valores do campo p1 [p1 está sendo passado por referência]
}
