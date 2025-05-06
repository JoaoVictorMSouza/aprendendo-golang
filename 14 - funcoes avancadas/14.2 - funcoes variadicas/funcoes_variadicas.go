package main

import "fmt"

// Usa-se o operador ... para indicar que a função pode receber um
// número variável de argumentos
// O tipo do argumento deve ser um slice, ou seja, um array de
// tamanho variável
func soma(numeros ...int) {
	fmt.Println("Números:", numeros)
}

func soma2(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	// ***
	// Funções variádicas
	// ***
	// São funções que podem receber um número variável de argumentos
	fmt.Println("Funções variádicas")

	soma()
	fmt.Println("------------------------------------")

	soma(1, 2, 3, 4, 5)
	fmt.Println("------------------------------------")

	soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("------------------------------------")

	somaTotal := soma2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Soma total:", somaTotal)
	fmt.Println("------------------------------------")

	escrever("Número", 1, 2, 3, 4, 5)
	fmt.Println("------------------------------------")
}
