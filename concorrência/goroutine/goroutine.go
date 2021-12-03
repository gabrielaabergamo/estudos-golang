package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//processo serial: primeiro executa a função da linha 17 e depois a da linha 18
	// fale("Squirtle", "squirtlesquirtlee", 3)
	// fale("Bulbasaur", "bulbabulba", 1)

	//vai executar de forma independente
	go fale("Squirtle", "squirtlesquirtlee", 500)
	go fale("Bulbasaur", "bulbabulba", 500)
	time.Sleep(time.Second * 5)
	fmt.Println("fimmmmm")
}
