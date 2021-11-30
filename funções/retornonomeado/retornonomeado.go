package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	primeiro = p1
	segundo = p2
	return //retorno limpo
}

func main() {
	elem1, elem2 := trocar(2, 10)
	fmt.Println(elem1, elem2)

	segundo, primeiro := trocar(2, 10)
	fmt.Println(segundo, primeiro)
}
