package main

import "fmt"

func main() {
	a, b := 230, 27
	var soma= a + b
	var sub= a - b
	var mult= a * b
	var div = float64(a) / float64(b)

	fmt.Println("valores:")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("soma = ", soma, "\nsubtração = ", sub, "\nMultiplicação= ", mult, "\ndivisão= ",div )
}
