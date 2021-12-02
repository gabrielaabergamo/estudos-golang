package main

import (
	"fmt"
	"reflect"
)

type curso struct {
	nome string
}

func main() {
	//parecido com o Object em java
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "opa"
	fmt.Println(coisa2, reflect.TypeOf(coisa2))

	coisa2 = curso{"beerus deus da destruição"}
	fmt.Println(coisa2, reflect.TypeOf(coisa2))
}
