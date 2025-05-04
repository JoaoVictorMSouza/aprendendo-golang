package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	i := 0

	// ***
	// Loop simples com for
	// ***
	// O for é o único loop em Go, não existe while ou do while
	for i < 5 {
		fmt.Println(i)
		i++
		time.Sleep(time.Second)
	}

	fmt.Println(i)
	fmt.Println("------------------------------------")

	// ***
	// Loop com for iniciando variável, condição e incremento
	// ***
	for j := 0; j < 5; j += 2 {
		fmt.Println(j)
		time.Sleep(time.Second)
	}
	fmt.Println("------------------------------------")

	// ***
	// Loop para iterar sobre um array, slice, string ou map
	// O range retorna o índice e o valor do elemento
	// ***

	// Loop com array
	var nomes [5]string = [5]string{"João", "Maria", "José", "Ana", "Pedro"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
		time.Sleep(time.Second)
	}

	fmt.Println("------------------------------------")

	// Loop com slice
	var nomes2 []string = []string{"João", "Maria"}
	for _, nome := range nomes2 {
		fmt.Println(nome)
		time.Sleep(time.Second)
	}

	fmt.Println("------------------------------------")

	// Loop com string
	for i, letra := range "Lucas" {
		fmt.Println(i, letra, string(letra))
		time.Sleep(time.Second)
	}

	fmt.Println("------------------------------------")

	// Loop com map
	usuario := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
		time.Sleep(time.Second)
	}
	// Não é possível loops em Structs
}
