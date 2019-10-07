package main

import "fmt"

func main() {
	//
	//ex1()
	//ex2()
	//ex3()
	ex4()
}

func ex1()  {
	number := 55

	if number%2 == 0 {
		fmt.Printf("O número %v é Par!", number)
	} else {
		fmt.Printf("O número %v é IMPar!", number)
	}
}

func ex2()  {
	birthYean := 1980

	yearNow := 2019

	age := yearNow - birthYean

	if yearNow >= birthYean {
		if age >= 16 {
			fmt.Printf("essa pessoa tem %v anos, Pode votar!\n", age)
		} else if age < 16 && age >= 0 {
			fmt.Printf("essa pessoa tem %v anos, Não podi vota.\n", age)
		}
	} else {
		fmt.Println("Valores inválidos, tente de novo.")
	}
}

func ex3()  {

	var senha string = "naotemsenha"
	entrada := "naotemsenha"
	//entrada := "qlq outra"

	if (senha == entrada){
		fmt.Println("ACESSO PERMITIDO!")
	} else {
		fmt.Println("ACESSO NEGADO")
	}
	
}

func ex4()  {
	variavel := false

	tipoVariavel := fmt.Sprintf("%T", variavel)

	if (tipoVariavel == "string"){
		fmt.Println( variavel,"> É uma string")
	} else if tipoVariavel == "int"{
		fmt.Println(variavel, "> É um número inteiro")
	} else if tipoVariavel == "bool"{
		fmt.Println(variavel, "> É um boolean")
	} else if tipoVariavel == "float64"{
		fmt.Println(variavel, "> É um float")
	} else {
		fmt.Println("Tipo da variável não identificado!")
	}
}