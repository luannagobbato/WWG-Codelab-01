package main

import "fmt"

func main() {
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