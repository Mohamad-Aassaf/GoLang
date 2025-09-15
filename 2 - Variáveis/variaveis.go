package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1" // declaração explícita de tipo
	variavel2 := "Variavel 2" // inferência de tipo, o go identifica o tipo da variável automaticamente

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6" // declaração de múltiplas variáveis
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1" // declaração de constante, o valor não pode ser alterado
	constante2 := "Constante 2"
	fmt.Println(constante1, constante2)

	variavel5, variavel6 = variavel6, variavel5 // trocando valores de variáveis
	fmt.Println("Troca de Valor de Variaveis::: ", variavel5, variavel6)
}