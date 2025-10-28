package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	nomero := 10

	if nomero < 0 {
		fmt.Println("Número negativo")
	} else if nomero > 0 {
		fmt.Println("Número positivo")
	} else {
		fmt.Println("Número é zero")
	}
}