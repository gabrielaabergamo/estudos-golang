package main

import "fmt"

func main() {
	s := make([]int, 10) //criamos um slice sem basear em nenhum array, foi criado um array interno
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) //cirando array interno com 20 elementos porem o slice contempla 10
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s)

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
