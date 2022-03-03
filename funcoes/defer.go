package main

import "fmt"

//Estratégia utilizada para executar uma determinada setença no último momento possivel dentro da função
func obtemValorAprovado(num int) int {
	defer fmt.Println("Fim!!!")
	if num > 5000 {
		fmt.Println("Valor alto!!!")
		return 5000
	}
	fmt.Println("Valor baixo!!!")
	return num
}
func main() {
	fmt.Println(obtemValorAprovado(6000))
	fmt.Println(obtemValorAprovado(3000))
}
