package main

import "fmt"

func somar(n1 int8, n2 int8) int8 { //Dentro do parenteses, variaveis da função, fora do parenteses = oque será retornado
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) { // declara tipos de retorno (2)
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func() { //Função é um tipo
		fmt.Println("Função F")
	}

	f() // chamar função f

	var g = func(txt string) { //FOutra forma de declarar
		fmt.Println(txt)
	}

	// resultado := g("Texto da função g") // chamar função g com parametro
	// fmt.Println(resultado)

	g("Texto da função g")
	

	resultadosSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadosSoma)
}
