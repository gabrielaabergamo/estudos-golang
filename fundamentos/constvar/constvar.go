package main

import (
	"fmt"
	"math"
	// _ "error": com o _ estamos importando error mas nn vamos utilizá-lo
	// m "math": vamos utilizar o m para referenciar Math-> em vez de Math.pow() temos m.pow()
)

func main() {
	//sempre que definirmos uma variável precisamos utilizá-la obrigatoriamente
	const PI float64 = 3.14
	var raio = 3.2 //tipo float64 inferido pelo compilador

	//forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é: ", area)

	// outra forma de criar constantes
	const (
		a = 1
		b = 2
	)

	// outra forma de criar variáveis
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	//criando duas variáveis em uma mesma linha de código
	var e, f bool = true, false
	fmt.Println(e, f)

	//também podemos fazer por inferência
	g, h, i := 2, false, "epa"
	fmt.Println(g, h, i)
}
