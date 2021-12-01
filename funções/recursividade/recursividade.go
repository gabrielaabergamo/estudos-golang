package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	aux, _ := fatorial(5)
	fmt.Println(aux)

	_, erro := fatorial(-4)
	if erro != nil {
		fmt.Println(erro)
	}
}
