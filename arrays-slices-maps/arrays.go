package main

import "fmt"

func main() {

	//Arrays
	var notas [3]float64
	fmt.Println(notas)
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.6
	fmt.Println(notas)

	fmt.Println("###########################")

	numeros := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // com essa inicialização, compilador conta e infere o tamanho do array

	for i, numero := range numeros {
		fmt.Printf("[%d]: %d\n", i, numero)
	}

}
