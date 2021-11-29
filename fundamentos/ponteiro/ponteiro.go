package main

import "fmt"

func main() {
	i := 1

	var p *int = nil //ponteiro apontando para nulo
	p = &i           // atribuindo o endereço da variável i para o ponteiro p
	*p++             //acessando o valor contido no endereço da variável que foi atribuída à p e incrementando
	i++

	//não é possível fazer aritmética de ponteiros
	//p++ não pode ser feito

	fmt.Println(p, *p, i, &i)

}
