package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64 := 10, 20, 30, 40 // intNumero é o tamanho do inteiro em bits, se não especificar o go usa a arquitetura da máquina (32 ou 64 bits)
	var numero int16 = 100
	fmt.Println(numero)

	// uint // unsygned int, ou seja, não aceita números negativos
	var numero2 uint = 1000 // só aceita números positivos, se tentar colocar número negativo dá erro de compilação e tem uint16, uint32, uint64 etc
	fmt.Println(numero2)

	// alias
	var numero3 rune = 12456 // INT32 = RUNE
	fmt.Println(numero3)

	var numero4 byte = 123 // UINT8 = BYTE
	fmt.Println(numero4)

	// NUMEROS REAIS (float32, float64)
	var numeroReal1 float32 = 123.45 //sempre ponto, nunca vírgula
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000000000000.45 // mais usado pq tem mais precisão
	fmt.Println(numeroReal2)

	// infererência de tipo
	numeroReal3 := 123.4567890123456789 // o go infere que é float64 pq tem mais precisão
	fmt.Println(numeroReal3)

	// STRING

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'A' // char é um int32 (rune), representa o valor da tabela ascII
	fmt.Println(char)
	fmt.Printf("%T\n", char) // %T mostra o tipo da variável

	// FIM STRINGS

	var texto int
	fmt.Println(texto) //valor padrão é vazio, string = linha vazia, int = 0, float = 0.0, bool = false

	// BOOLEAN
	var booleano1 bool = true
	fmt.Println(booleano1)

	booleano2 := false
	fmt.Println(booleano2)

	// VALOR ERRO

		// nome da variavel, tipo, valor
	var erro error = errors.New("Erro personalizado") // valor padrão é nil (vazio)
	fmt.Println(erro)
}