package main

import "fmt"

//func nome_função(lista de parâmetros) (tipo/tipos de retorno) {...}

func f1() {
	fmt.Println("F1")
}

func f2(p1, p2 string) {
	fmt.Println(p1, p2)
}

func f3() string {
	return "F3"
}

func f5() (string, string) {
	return "Retorno1", "Retorno2"
}

func main() {
	f1()
	f2("Pokémon", "é legal")
	fmt.Println(f3())
	retorno1, retorno2 := f5()
	fmt.Println(retorno1, retorno2)
	fmt.Println(f5())
}
