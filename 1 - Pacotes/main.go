package main

import (
	"fmt"
	"modulo/auxiliar"
	
	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("Escrevendo o arquivo main!")
	auxiliar.Escrever() // chamando a função do outro pacote, se uma função começa com letra maiuscula ela é publica, se começa com minuscula é privada
	// auxiliar.escrever2() não funciona pq a função começa com letra minuscula, ou seja, é privada e só pode ser usada dentro do pacote auxiliar
	
	erro := checkmail.ValidateFormat("emailLegal@gmail.com") // valida formato de um email

	fmt.Println(erro)
	fmt.Println(checkmail.ValidateFormat("emailErrado!gmail.com"))
}

// go build <- comando pra compilar o código em executável, modulo.exe executa no terminal'
// go run /1 - Pacotes/main.go <- comando para rodar o arquivo