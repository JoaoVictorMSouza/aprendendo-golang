package main

import "fmt"

// ***
// Função comum parâmetro e retorno
// ***
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// ***
// Função com vários retornos
// ***
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

// ***
// Armazenando uma função em uma variavel
// ***
var funcao = func(texto string) string {
	fmt.Println("Essa é uma função anônima")
	fmt.Println(texto)
	return texto
}

func main() {
	resultado := somar(5, 10)
	fmt.Println("Resultado da soma:", resultado)

	resultadoFuncaoAnonica := funcao("Texto da função anônima")
	fmt.Println("Resultado da função anônima:", resultadoFuncaoAnonica)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 5)
	fmt.Println("Resultados dos cálculos:", resultadoSoma, " | ", resultadoSubtracao)

	_, resultadoSubtracao2 := calculosMatematicos(10, 5)
	fmt.Println("Resultado da soma:", resultadoSubtracao2)
}
