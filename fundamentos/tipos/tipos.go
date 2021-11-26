package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255 // byte = alias de uint8
	fmt.Println("O tipo de byte é", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo de int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representacao da tabela unicode (int32)
	fmt.Println("O tipo de rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	//números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo de 49.99 é", reflect.TypeOf(49.99)) //float64

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))

	//string
	s1 := "Salve fellas"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho de s1 é", len(s1))

	//string com múltiplas linhas
	s2 := `Salve
	fellas`
	fmt.Println(s2)
	fmt.Println("O tamanho de s2 é", len(s2))

	//char???
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
	//var x char = 'b' não existe
}
