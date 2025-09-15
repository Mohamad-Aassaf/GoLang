package auxiliar

import "fmt"

// Escrever registra uma mensagem na tela
func Escrever() { //  se uma função começa com letra maiuscula ela é publica, se começa com minuscula é privada
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2() // chamando a função privada dentro do mesmo pacote
}