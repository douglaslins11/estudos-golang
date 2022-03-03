package main

import "fmt"

type Funcionario struct {
	nome    string // Campo não é visível fora do pacote
	Idade   int    //Campo exportado e visível fora do pacote
	salario float32
}

type Endereco struct {
	rua    string
	cidade string
	estado string
}

type Pessoa struct {
	nome     string
	endereco Endereco
}

func main() {
	func1 := Funcionario{
		nome:    "Douglas",
		Idade:   25,
		salario: 1200.25,
	}

	func2 := Funcionario{"Douglas", 25, 1200.25}

	fmt.Println("Func 1:", func1)
	fmt.Println("Func 2:", func2)

	//Acessando campos específicos
	fmt.Println("Func1 - Nome:", func1.nome)

	//Structs aninhadas
	pess := Pessoa{
		nome: "Ney Silva",
		endereco: Endereco{
			rua:    "Rua 123",
			cidade: "Recife",
			estado: "PE",
		},
	}

	fmt.Println("Pessoa:", pess)
}
