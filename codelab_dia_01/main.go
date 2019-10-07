package main

import "fmt"

func main() {
	//##Exercício 1

	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
}

func ex1(){
	var a = 10
	var b  = 15
	var c = 25

	fmt.Println("#Excercício 1")
	fmt.Println("a=", a , "\tb=", b ,"\tc=", c)
}

func ex2(){
	a,b := 230, 27
	var soma= a + b
	var sub= a - b

	fmt.Println("#Excercício 2")
	fmt.Println("valores:")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("soma = ", soma, "\nsubtração = ", sub)
}

func ex3(){
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

func ex4() {
	a, b := 230, 27
	var soma= a + b
	var sub= a - b
	var mult= a * b
	var div = float64(a) / float64(b)

	fmt.Printf("Tipos das variáveis:\na= %T,\nb= %T,\nsoma= %T,\nsub= %T,\nmult= %T,\ndiv= %T", a,b,soma,sub,mult,div)
}

func ex5(){
	//#04
	//Preços dos itens do mercado

	var preçoDoPão float64 = 4.60
	var preçoDaAveia float64 = 5
	var preçoDoAzeite float64 = 19.95
	var preçoDoSuco float64 = 7
	var preçoDaÁgua float64 = 2.15
	var preçoDoKGMaçã float64 = 8.90
	var preçoDoKGBanana float64 = 4.5

	/*- 3 pães
		- 1 azeite
	- 2 garrafas de suco de laranja
	- 5 garrafas de água
	- 1.5 kg de maçã
	*/

	custoCompra := preçoDoPão * 3 +
		preçoDoAzeite * 1 +
		preçoDoSuco * 2 +
		preçoDaÁgua * 5 +
		preçoDoKGMaçã * 1.5 +
		preçoDoKGBanana * 0 +
		preçoDaAveia * 0

	fmt.Println("O preço total da compra foi: R$",custoCompra )
}

func ex6(){
	// var := 27
	// fmt.Println("A= ", var)
	fmt.Println("A palavra `var` é uma palavra reservada da linguagem, não podendo ser" +
		"utilizada como nome de variável.")
	// var é uma palavra reservada da linguagem go e
}

func ex7(){
	var a = 15
	b:= 31
	var c int = 47

	fmt.Printf("Valor Original: %v\t%v\t%v", a, b, c)
	fmt.Printf("\nBase 2: %b\t%b\t%b", a, b, c)
	fmt.Printf("\nBase 10: %d\t%d\t%d", a, b, c)
	fmt.Printf("\nBase 16: %X\t%X\t%X", a, b, c)
}

func ex8()  {
	var x int
	var f bool
	var str []string

	fmt.Println("X = ", x, "F = ", f, "STR =", str)
	// Em Go todas as variaveis possuem um valor default, assim, elas sempre são inicializadas com eles quando não possuem valor atribuido
}


