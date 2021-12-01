package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	aux1 := fatorial(5)
	fmt.Println(aux1)

	//aux2 := fatorial(-4) gera problema de tipo (-4 não é uint)
}
