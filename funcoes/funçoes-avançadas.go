package main

import "fmt"

//Armazenar funções em variáveis
var soma = func(a, b int) int { return a + b }

func multiplicacao(a, b int) int { return a * b }

func exec(funcao func(int, int) int, param1, param2 int) int {
	return funcao(param1, param2)
}

func main() {
	fmt.Println(soma(3, 4))
	fmt.Println(exec(multiplicacao, 4, 3))
}
