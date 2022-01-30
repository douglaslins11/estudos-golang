package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Slice é uma estrutura de dados que tem um ponteiro apontando para um array e um tamanho

	a1 := [3]int{1, 2, 3} //array (tamanho único)
	s1 := []int{1, 2, 3}  //slice (tamanho variável)

	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	//Slice não é um array!! Ele define um pedaço/parte de um array.
	a2 := [5]int{1, 2, 3, 4, 5}

	s2 := a2[1:3] //Vai do indice 1 ao 3 do array acima, sem incluir o indice 3
	fmt.Println(s2)

	//Novo Slice, porém apontando para o mesmo array
	s3 := a2[:2] //Vai inicio ao indice 2, sem incluir o indice 2
	fmt.Println(s3)

	s4 := make([]int, 10)
	s4[9] = 12
	fmt.Println(s4)

	s4 = make([]int, 10, 20)

	/* Na criação do slice s4(linha 27), foi criado um Slice de inteiros com 10 posições, por baixo dos panos o compilador
	cria um array também com  posições, para onde o slice aponta.
	   Na linha 31, foi criado um novo slice, porém, o terceiro parâmetro da função especifica o tamanho do array, então
	nesse caso, foi criado um array de 20 elementos e um slice de 10 elementos, que aponta para esse array
	*/

	fmt.Println("########################")
	fmt.Println("Características do Slice:")
	fmt.Printf("Slice: %d\n", s4)
	fmt.Printf("Tamanho: %d\n", len(s4))
	fmt.Printf("Capacidade máxima: %d\n", cap(s4))

	s4 = append(s4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s4, len(s4), cap(s4)) // Capacidade máxima do slice atingida
	s4 = append(s4, 1)
	/*A execução da linha acima não dar erro, mesmo extrapolando o tamanho máximo do slice,
	internamente, ele aponta para outro array com a cópia dos elementos do array atual e
	o dobro do tamanho*/

	/*No exemplo abaixo conseguimos criar 2 slices apontando para o mesmo array, isso é provado com
	a alteração feita na posição 0 no slice s5(linha 59) e automaticamente é refletida no slice s6,
	pois internamente a alteração foi feira no array e os 2 apontam para o mesmo

	Essa é a vantagem do slice, conseguimos trabalhar com uma estrutura de dados super leve e flexível,
	com a performance de uma estrutura de dados estática, devido ao fato de ter uma por baixo.*/
	s5 := make([]int, 10, 20)
	s6 := append(s5, 1, 2)
	fmt.Println(s5, s6)

	s5[0] = 7
	fmt.Println(s5, s6)
}
