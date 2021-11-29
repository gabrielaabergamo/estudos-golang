package main

import (
	"fmt"
	"reflect"
)

func main() {
	//slice não é um array! é um pedaço de um array
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf((a1)), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3] //slice do índice 1 até o 3 sem contar o 3
	fmt.Println(a2, s2)
	//slice não gera um array novo! é uma estrutura que tem um ponteiro que aponta para elementos de um determinado array
	s3 := a2[:2]
	fmt.Println(a2, s3)

}
