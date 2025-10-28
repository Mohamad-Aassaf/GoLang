package main

import "fmt"

func main()  {
	fmt.Println("Arrays e Slices")

	var array1 [5]string //se não colocar nada, vai dar tudo 0, nulo, branco etc
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // [...] deixa o tamanho com o valor baseado na quantidade de valores passados no{}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18) // pega o slice, add mais um item e retorna um slice novo
	fmt.Println(slice)

	slice2 := array2[1:3] // inclusivo : exclusivo
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	// 9.1 Arrays Internos
	fmt.Println("Arrays Internos")
	slice3 := make([]float32, 10, 11) // cria um slice de float32 com 10 posições e capacidade pra 15
	fmt.Println(slice3)

	slice3 = append(slice3, 5) // adiciona mais um elemento no slice
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // len = tamanho
	fmt.Println(cap(slice3)) // cap = capacidade

	slice4 := make([]float32, 5) // se não colocar a capacidade, ela vai ser igual ao tamanho
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}