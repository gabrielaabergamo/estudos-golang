package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//maps precisam ser inicializados

	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[46376473647] = "Pedro"
	aprovados[47387483785] = "Ana"
	fmt.Println(aprovados)

	for cpf, pessoa := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", pessoa, cpf)
	}

	fmt.Println(aprovados[12345678978])
	delete(aprovados, 12345678978)
	fmt.Println(aprovados)
}
