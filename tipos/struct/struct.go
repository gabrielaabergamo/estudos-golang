package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//método: função com receiver
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}
func main() {
	var prod produto
	prod = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(prod, prod.precoComDesconto())

	prod2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(prod2.precoComDesconto())
}
