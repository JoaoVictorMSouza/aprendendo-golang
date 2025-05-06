package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var varriavel2 int = variavel1

	fmt.Println(variavel1, varriavel2)

	variavel1++
	fmt.Println(variavel1, varriavel2)
	fmt.Println("------------------------------------")

	// Ponteiros são variáveis que guardam o endereço
	// de outra variável na memória
	var variavel3 = 100
	var ponteiro *int
	fmt.Println(variavel3, ponteiro)
	fmt.Println("------------------------------------")

	// & é o operador que retorna o endereço de memória
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)
	fmt.Println("------------------------------------")

	// Declaração e inicialização do ponteiro
	var ponteiro2 *int = &variavel3
	fmt.Println(variavel3, ponteiro2)
	fmt.Println("------------------------------------")

	// * é o operador que retorna o valor armazenado no endereço de memória
	fmt.Println(*ponteiro2)
	fmt.Println("------------------------------------")
	variavel3++
	fmt.Println(variavel3, ponteiro2)
	fmt.Println(*ponteiro2)
	fmt.Println("------------------------------------")
}
