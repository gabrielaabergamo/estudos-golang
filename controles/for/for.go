package main

import "fmt"

func main() {
	//não existe while em go
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	//laço infinito: for{...}
}
