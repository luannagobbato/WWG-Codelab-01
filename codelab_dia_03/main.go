package main

import "fmt"

func main() {
	a, b := 230, 27
	var soma = a + b
	var sub = a - b
	var mult = a * b
	var div = float64(a) / float64(b)

	fmt.Println("valores:")
	fmt.Println("a =", a, "\nb =", b)
	fmt.Println("soma = ", soma, "\nsubtração =", sub)
	fmt.Println("multiplicação =", mult, "\ndivisão =", div)

	// #3
	fmt.Printf("\nTipos\na: %T\nb: %T\nSoma: %T\n Subtração: %T\nDivisão: %T\nMultiplicação: %T", a,b,soma,sub,div,mult)
}