package main

import "fmt"

type usuario struct { // tipo - nome - struct | Struct = tipo uma classe
	nome     string
	idade    uint8
	endereco endereco // variavel endereco do tipo endereco, struct dentro de struct
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario //struct sem passar valor da valor 0 para todos
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bolsonaros", 0}

	usuario2 := usuario{"Cleber", 21, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Diogenes"} // ou idade: 22
	fmt.Println(usuario3)
}
