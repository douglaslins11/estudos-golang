package main

import "fmt"

type Esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {

	carro1 := ferrari{modelo: "F40", turboLigado: false, velocidadeAtual: 0}
	carro1.ligarTurbo()

	var carro2 Esportivo = &ferrari{modelo: "F40", turboLigado: true, velocidadeAtual: 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
