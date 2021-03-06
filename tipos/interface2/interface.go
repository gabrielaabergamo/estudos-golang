package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() { //interface alterando struct
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"F40", false, 0} //trabalhando no nível de interface
	//sem o &: interface.go:23:6: cannot use ferrari{...} (type ferrari) as type esportivo in assignment:
	//ferrari does not implement esportivo (ligarTurbo method has pointer receiver)
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
