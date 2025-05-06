package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Recuperando a execução")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 && media > 0 {
		return false
	}

	panic("Nota menor ou igual a zero")
}

func main() {
	// ***
	// Panic e recover
	// ***
	// O panic é uma forma de interromper a execução normal do
	// programa e iniciar o processo de recuperação.
	// O recover é uma função que pode ser usada para recuperar
	// de um panic.
	fmt.Println("Panic e recover")

	fmt.Println(alunoEstaAprovado(7, 8))
	fmt.Println("------------------------------------")

	fmt.Println(alunoEstaAprovado(0, 0))
	fmt.Println("------------------------------------")
	fmt.Println("FIM")
}
