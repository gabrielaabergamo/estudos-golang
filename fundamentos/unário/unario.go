package main

import "fmt"

func main() {
	x := 1
	y := 2

	x++ //postfix
	// ++x não existe em go
	fmt.Println(x)

	y-- //postfix
	fmt.Println(y)

	// (x == y--) não pode
}
