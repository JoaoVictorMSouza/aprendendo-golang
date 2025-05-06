package main

import "fmt"

func main() {
	// ***
	// Funções anônimas
	// ***
	// São funções que não possuem nome e podem ser atribuídas a variáveis
	// ou passadas como argumentos para outras funções
	fmt.Println("Funções anônimas")

	// ***
	// Função anônima sem parâmetro e sem retorno
	// ***
	// A função está sendo declarada e logo em seguida sendo invocada pelos ()
	func() {
		fmt.Println("Essa é uma função anônima")
		fmt.Println("------------------------------------")
	}()

	// ***
	// Função anônima com parâmetro e sem retorno
	// ***
	func(texto string) {
		fmt.Println(texto)
		fmt.Println("------------------------------------")
	}("Essa é uma função anônima com parâmetro")

	// ***
	// Função anônima com parâmetro e retorno
	// ***
	retorno := func(texto string) string {
		fmt.Println(texto)
		fmt.Println("------------------------------------")
		return fmt.Sprintf("Recebido: %s", texto)
	}("Essa é uma função anônima com parâmetro e retorno")
	fmt.Println(retorno)
}
