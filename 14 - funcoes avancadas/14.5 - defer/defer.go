package main

import "fmt"

func funcao1() {
	fmt.Println("Função 1")
	fmt.Println("------------------------------------")
}

func funcao2() {
	fmt.Println("Função 2")
	fmt.Println("------------------------------------")
}

func funcao3() {
	fmt.Println("Função 3")
	fmt.Println("------------------------------------")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	media := (n1 + n2) / 2

	// Com isso, somente será printado no console imediatamente
	// antes do return
	defer fmt.Println("Media calculada:", media)

	fmt.Println("Aluno está aprovado?")

	if media >= 7 {
		// Repetição de código, não é uma boa prática, por isso o defer
		// fmt.Println("Media calculada:", media)
		return true
	}

	// Repetição de código, não é uma boa prática, por isso o defer
	// fmt.Println("Media calculada:", media)
	return false
}

func main() {
	// ***
	// Defer
	// ***
	// O defer é usado para adiar a execução de uma função até
	// o final da função atual
	// Ele é útil para garantir que uma função seja executada
	// DEFER = ADIAR
	fmt.Println("Defer")

	defer funcao1()
	funcao2()
	funcao3()

	fmt.Println(alunoEstaAprovado(7, 6))
	fmt.Println("------------------------------------")
}
