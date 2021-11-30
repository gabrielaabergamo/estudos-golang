package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d. %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Tiago", "Guilherme"}
	imprimirAprovados(aprovados...)

	//aprovados := [4]string{"Maria", "Pedro", "Tiago", "Guilherme"} não daria certo
	//aprovados := [...]string{"Maria", "Pedro", "Tiago", "Guilherme"} não daria certo
}
