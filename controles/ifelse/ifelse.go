package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}

	//obs: 1. não tem () no if [if(x)], o certo é if x [porém se for uma condição complexa aí pode usar ()]
	// if nota >= 7 && (true || false) está correto
	//2. não podemos fazer if sem {}, logo if x return 1 está errado, o correto é if x { return 1 }
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
