package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	fmt.Println("Teste " + string(97)) // string(97) = a

	//int para string
	fmt.Println("Teste " + strconv.Itoa(97))

	//string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("1")
	if b {
		fmt.Println("Truuuue")
	}
}
