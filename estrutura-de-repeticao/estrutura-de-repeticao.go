package main

import "fmt"

func main() {
	//A única estrutura de repetição presente na linguagem é o "for"

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i<=20; i+=2 {
		fmt.Println(i)
	}

	for {
		//loop infinito
	}


}
