package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	//inteiros
	fmt.Println("########## Inteiros")
	fmt.Println(1,2,10000)
	fmt.Println("Literal inteiro: ", reflect.TypeOf(20000))

	//sem sinal(apenas positivos)... uint8 uint16 uint32 uint64
	var a byte=255
	fmt.Println("Byte: ", reflect.TypeOf(a))

	//com sinal.. int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("Valor max inteiro: ", i1)
	fmt.Println("Typo i1: ", reflect.TypeOf(i1))

	//representa um mapeamento da tabela Unicode (int32)
	var i2 rune = 'a'
	fmt.Println("O rune é: ", reflect.TypeOf(i2))
	fmt.Println("i2: ", i2)

	//números reais (float32, float64)
	fmt.Println("########## Reais")
	var flt float32 = 49.99
	fmt.Println("O tipo de f: ", flt)
	fmt.Println("O tipo do literal 49.99 é: ", reflect.TypeOf(49.99)) //float64

	//Boolean
	var t bool = true
	var f bool = false
	fmt.Println(t,f)

	//String
	s1 := "Douglas Lins da Silva"
	s2 := `Douglas
Lins
da
Silva`

	fmt.Println(s1)
	fmt.Println(s2)


	//Tipos Default
	var INTEIRO int
	var REAL float64
	var BOLEANO bool
	var STRING string
	var PONTEIRO *int
	fmt.Println("########## Tipos Default")
	fmt.Println("Default INT: ", INTEIRO)
	fmt.Println("Tipos FLOAT", REAL)
	fmt.Println("Tipos BOOLEAN", BOLEANO)
	fmt.Println("Tipos STRING", STRING)
	fmt.Println("Tipos PONTEIRO", PONTEIRO)
}
