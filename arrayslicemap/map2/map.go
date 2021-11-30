package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro":          1200.0,
	}

	funcsESalarios["Tiago Filho"] = 1350.0 //adicionando elem. ao map
	delete(funcsESalarios, "Inexistente")  //tentar excluir um elemento que não existe não gera problemas

	for nome, salário := range funcsESalarios {
		fmt.Println(nome, salário)
	}
}
