package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //compilador que vai contar o número de elementos

	for i, numero := range numeros { //range numeros vai retornar o i e o numeros[i]
		fmt.Printf("%d. %d\n", i, numero)
	}
	fmt.Println(" ")

	for _, num := range numeros {
		fmt.Println(num)
	}
	fmt.Println(" ")

	for indice := range numeros {
		fmt.Println(indice)
	}
}
