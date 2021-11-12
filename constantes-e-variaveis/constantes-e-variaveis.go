package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float 64) inferido pelo compilador, tipagem dinamica

	//forma reduzida de criar uma variável, ela já é  declarada e inicializada
	//caso uma variável seja declada e não utilizada, acontece um erro de compilação
	area := PI * math.Pow(raio,2)
	fmt.Println("Area: ", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a,b,c,d)

	var e, f bool = true, false
	fmt.Println(e,f)

	g, h, i := 2, false, "Oxe!"
	println(g,h,i)
}