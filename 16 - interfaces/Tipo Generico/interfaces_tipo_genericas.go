package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	// ***
	// Interfaces genéricas
	// ***
	// Interfaces genéricas são interfaces que podem ser usadas
	// com diferentes tipos, sem a necessidade de definir uma
	// interface específica para cada tipo. Elas permitem que
	// funções e métodos sejam escritos de forma mais flexível,
	// aceitando qualquer tipo que implemente os métodos
	// definidos na interface. Isso é útil para criar funções
	// e métodos que podem trabalhar com diferentes tipos de dados
	// de forma genérica, sem a necessidade de duplicar código
	// para cada tipo específico.

	generica("Olá, mundo!")
	generica(42)
	generica(3.14)
	generica([]string{"Go", "é", "legal"})
	generica(map[string]int{"um": 1, "dois": 2, "três": 3})
	generica(true)
}
