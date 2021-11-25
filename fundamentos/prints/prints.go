package main

import "fmt"

func main() {
	fmt.Print("Mesma ")
	fmt.Print("linha.")

	fmt.Print("\n")

	fmt.Println("Nova")
	fmt.Println("linha.")

	x := 3.14
	xs := fmt.Sprint(x)

	fmt.Println("O valor é: " + xs)
	fmt.Println("O valor é:", x)
	fmt.Printf("O valor é: %f", x)

	a := 1
	b := 1.999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
