package main

import "fmt"

/*Método é uma função com um tipo de receptor especial
entre a palavra chave 'func' e o nome do método
Ex: func (t Type) nomeDoMetodo (parametros) {
}
*/

type Funcionario struct {
	nome    string
	moeda   string
	salario int
}

func (funcionario Funcionario) exibeSalario() {
	fmt.Printf("O salário de %s é %s%d\n", funcionario.nome, funcionario.moeda, funcionario.salario)
}

func (funcionario *Funcionario) alteraSalario() {
	funcionario.salario = 1234
}

func main() {
	func1 := Funcionario{"Douglas", "$", 5000}
	func1.exibeSalario()
	func1.alteraSalario()
	func1.exibeSalario()
}
