package main

import (
	"fmt"
)

func fatorial(numero int) int {
	if numero == 0 {
		return 1
	}

	return fatorial(numero-1) * numero
}

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	// ***
	// Funções recursivas
	// ***
	// São funções que se chamam dentro delas mesmas
	fmt.Println("Funções recursivas")

	retornoFatorial := fatorial(3)
	fmt.Println("Fatorial de 3:", retornoFatorial)
	fmt.Println("------------------------------------")

	for index := uint(1); index < 5; index++ {
		fmt.Println("Fibonacci de", index, ":", fibonacci(index))
	}
}
