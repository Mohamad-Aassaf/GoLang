package main

import "fmt"

func main()  {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1 // uma copia da de cima vai pra de baixo, dessa linha pra baixo não muda mais nadd

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERENCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int // ponteiro de int, guarda o endereço de memoria de um valor inteiro

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) // pega o valor que ta no endereço


}