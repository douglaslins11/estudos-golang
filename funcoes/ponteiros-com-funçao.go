package main

import "fmt"

func incrementoPorValor(numero int) {
	numero++
}

func incrementroPorReferencia(numero *int) {
	*numero++
}

func main() {
	//Por padrão, em go, a passagem de parâmetro para funções, é por valor, caso queira passar por referência, tem que especificar no código

	num := 10

	incrementoPorValor(num) //valor não é alterado
	fmt.Println(num)

	incrementroPorReferencia(&num) //valor é alterado
	fmt.Println(num)
}
