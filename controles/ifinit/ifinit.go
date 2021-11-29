package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { //suportado no switch
		fmt.Println("Ganhouuuu")
	} else {
		fmt.Println("F")
	}

	//fmt.Println(i): da erro; i foi inicializada para o if, depois que ele rodou o i é descartado
}
