package main

import "fmt"

func main() {
	// Aritméticos
	// +, -, /,*, % <-- Resto da divisão

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25 // int32 = 25 Não pode fazer nada com dois valores de TIPOS DIFERENTES.
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// ATRIBUIÇÃO

	var variavel1 string = "String"
	variavel := "String2"
	fmt.Println(variavel1, variavel)

	// RELACIONAIS

	fmt.Println(1 > 2) // Vai dar false pq um não é maior do que dois né carai

	fmt.Println(1 >= 2) // maior igual que
	fmt.Println(1 == 2) // igual
	fmt.Println(1 <= 2) // menor igual que
	fmt.Println(1 > 2)  // maior que
	fmt.Println(1 < 2)  // menor que
	fmt.Println(1 != 2) // diferente de

	// LÓGICOS
	fmt.Println("-----------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // e
	fmt.Println(verdadeiro || falso) // ou
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// UNÁRIOS
	numero := 10
	// numero = numero + 1

	numero++     // numero = numero + 1
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // num = num - 20

	numero *= 3 // numero = numero * 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)

	var texto string

	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)

}
