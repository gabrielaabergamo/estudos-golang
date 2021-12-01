package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	//* é usado para desreferenciar: ter acesso ao valor do elem que o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) //por valor
	fmt.Println(n)

	inc2(&n)
	fmt.Println(n)
}
