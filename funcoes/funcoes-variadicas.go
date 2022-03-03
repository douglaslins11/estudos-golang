package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0

	for _, numero := range numeros {
		total += numero
	}
	return total / float64(len(numeros))
}

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Aprovados:")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s", i+1, aprovado)
	}
}

func main() {
	fmt.Printf("Media: %.2f", media(7.7, 8.8, 9.9))

	//pasando um slice como parâmetro
	aprovados := []string{"Douglas", "Ale", "Leo", "Anderson"}
	imprimirAprovados(aprovados...)
}
