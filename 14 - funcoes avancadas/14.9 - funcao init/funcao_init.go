package main

import "fmt"

var numero int

// Função init é por arquivo, diferente do main que é por pacote
func init() {
	fmt.Println("Executando a função init")
	numero = 10
}

func main() {
	// ***
	// Função init
	// ***
	// Função init é executada automaticamente
	// quando o pacote é importado. Ela é úteil para inicializar
	// variáveis ou executar código que deve ser executado antes
	// da execução do programa.
	fmt.Println("Funçao init")

	fmt.Println("Número:", numero)
}
