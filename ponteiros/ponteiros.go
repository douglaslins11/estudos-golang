package main

import "fmt"

func main() {
	i := 10

	/*Go não tem aritmética de ponteiros, o que podemos fazer é compartilhar o ponteiro com a referência
	para que várias váriáveis consigam apontar para o mesmo espaço na memória
	*/

	var p *int=nil
	p=&i //pegando o endereço na memória da variável i
	//p aponta para posição da memória
	//*p pega o valor contido na posição
	fmt.Println("Posição na memória que o ponteiro está apontando: ", p)
	fmt.Println("Valor contido na posição na memória apontada pelo ponteiro: ", *p)
}
