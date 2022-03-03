package main

import "fmt"

func trocar(param1, param2 int) (segundo, primeiro int) {
	segundo = param2
	primeiro = param1
	return //retorno limpo
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)
}
