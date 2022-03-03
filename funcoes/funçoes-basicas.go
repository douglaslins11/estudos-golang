package main

import "fmt"

func semParametrosSemRetorno() {
	fmt.Println("Função semParametrosSemRetorno")
}

func recebeParametrosSemRetorno(param1 string, param2 string) {
	fmt.Printf("Função recebeParametrosSemRetorno: %s - %s", param1, param2)
}

func semParametrosComRetorno() string {
	return "Função semParametrosComRetorno"
}

func recebeParametrosComRetorno(param1, param2 string) string {
	return fmt.Sprintf("Função recebeParametrosComRetorno: %s - %s", param1, param2)
}

func semParametrosComRetornoMultiplo() (string, string) {
	return "Douglas", "Lins"
}

func main() {
	semParametrosSemRetorno()
	recebeParametrosSemRetorno("Douglas", "Lins")
	fmt.Println(semParametrosComRetorno())
	fmt.Println(recebeParametrosComRetorno("Douglas", "Lins"))
	param1, param2 := semParametrosComRetornoMultiplo()
	fmt.Println(param1, param2)
}
