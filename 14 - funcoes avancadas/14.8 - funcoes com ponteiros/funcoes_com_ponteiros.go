package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	// ***
	// Funções com ponteiros
	// ***
	// Funções com ponteiros são funções que recebem ponteiros
	// como parâmetros. Isso significa que elas podem modificar
	// o valor da variável original, e não uma cópia dela. Isso
	// é útil quando queremos modificar o valor de uma variável
	fmt.Println("Funções com ponteiros")

	numero := 10
	numeroIntertido := inverterSinal(numero)
	fmt.Println("Número invertido:", numeroIntertido)
	fmt.Println("Número original:", numero)
	fmt.Println("------------------------------------")

	numeroPonteiro := 40
	fmt.Println("Número original:", numeroPonteiro)
	inverterSinalPonteiro(&numeroPonteiro)
	fmt.Println("Número invertido:", numeroPonteiro)
}
