package main

import "fmt"

func main() {
	fmt.Println("If e Else")

	var numero int = 15

	// ***
	// If simples
	// ***
	if numero > 15 {
		fmt.Println("O número é maior que 15")
	} else {
		fmt.Println("O número é menor ou igual a 15")
	}

	fmt.Println("------------------------------------")

	// ***
	// If init
	// ***
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println(outroNumero)
		fmt.Println("O número é maior que 0")
	} else if numero < -10 {
		fmt.Println("O número é menor que -10")
	} else {
		fmt.Println("O número é menor ou igual a 0")
	}
	fmt.Println("------------------------------------")
}
