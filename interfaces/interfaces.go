package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s: R$ %.2f", p.nome, p.preco)
}

func imprimir(imp imprimivel) {
	fmt.Println(imp.toString())
}

func main() {
	pess := pessoa{nome: "Douglas", sobrenome: "Lins"}
	pess.toString()
	imprimir(pess)

	prod := produto{nome: "Macbook", preco: 10000.00}
	prod.toString()
	imprimir(prod)

	var coisa imprimivel = pessoa{nome: "Marielly", sobrenome: "Lins"}
	coisa.toString()
	imprimir(coisa)

	coisa = produto{nome: "Monitor", preco: 800.00}
	coisa.toString()
	imprimir(coisa)

}
