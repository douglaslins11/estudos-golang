package main

import "fmt"

func main() {

	//declaração make(map[tipo da chave] tipo do valor)
	aprovados := make(map[int]string)

	aprovados[1231321651465] = "João"
	aprovados[1354681321684] = "Maria"
	aprovados[9895357411846] = "Messi"
	aprovados[9966842135656] = "CR7"

	fmt.Println(aprovados)

	//Percorrer
	for cpf, nome := range aprovados {
		fmt.Printf("[%d]: %s\n", cpf, nome)
	}

	//deletando
	//se não encontrar, não dar nenhum erro, segue o jogo
	delete(aprovados, 9895357411846)
	fmt.Println(aprovados)

	//inserindo
	aprovados[123465789] = "Magrão"
	fmt.Println(aprovados)

	//Atualizando
	aprovados[1231321651465] = "Durval"
	fmt.Println(aprovados)

	aprovados[12] = "Teste"
	fmt.Println(aprovados)
}
