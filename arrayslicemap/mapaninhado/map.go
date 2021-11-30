package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15769.70,
			"Guga Pereira":   1237.10,
		},
		"J": {
			"José João": 1135.80,
		},
		"P": {
			"Pedro": 1200.00,
		},
	}

	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "G")
	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
