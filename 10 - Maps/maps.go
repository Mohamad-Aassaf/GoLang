package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string { // map[TipoDaChave]TipoDoValor
		"nome":      "Cleber",
		"sobrenome": "da Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"]) // Acessar o valor da chave nome

	usuario2 := map[string]map[string]string { // map[TipoDaChave]TipoDoValor
		"nome": {
			"primeiro": "Cleber",
			"ultimo":   "da Silva",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Jaguarão",
		},
	}

	fmt.Println(usuario2)
	// fmt.Println(usuario2["nome"]["primeiro"])  Acessar o valor da chave primeiro dentro de nome
	delete(usuario2, "nome") // Deletar o campo nome do map usuario2
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string { // Adicionar um novo campo ao map usuario2
		"nome": "Aquário",
	}
	fmt.Println(usuario2)

}