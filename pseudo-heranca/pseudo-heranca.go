package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	turboLigado bool
	carro       //campo anonimo
}

func main() {
	//Usando campos anonimos, conseguimos simular o compoportamento de uma herança
	f := ferrari{}
	f.turboLigado = true
	f.velocidadeAtual = 0
	f.nome = "F40"

	fmt.Println(f)
}
