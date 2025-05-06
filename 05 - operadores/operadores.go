package main

import "fmt"

func main() {
	// ***
	// Operadores Aritméticos
	// +, -, *, /, %
	// ***
	soma := 10 + 5
	subtracao := 10 - 5
	multiplicacao := 10 * 5
	divisao := 10 / 5
	resto := 10 % 3

	fmt.Println(soma, subtracao, multiplicacao, divisao, resto)
	fmt.Println("------------------------------------")

	// ***
	// Operadores de Atribuição
	// =, +=, -=, *=, /=, %=
	// ***
	variavel1 := "String"
	var variavel2 string = "String"

	fmt.Println(variavel1, variavel2)
	fmt.Println("------------------------------------")
	// ...

	// ***
	// Operadores de Comparação
	// ==, !=, >, <, >=, <=
	// ***
	fmt.Println(10 == 10) // true
	fmt.Println(10 != 10) // false
	fmt.Println(10 > 5)   // true
	fmt.Println(10 < 5)   // false
	fmt.Println(10 >= 10) // true
	fmt.Println("------------------------------------")
	// ...

	// ***
	// Operadores Lógicos
	// && (E), || (OU), ! (NÃO)
	// ***
	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false
	fmt.Println("------------------------------------")

	// ***
	// Operadores unários
	// +, -, ++, --, -=, +=, *=, /=, %=
	// ***
	numero := 10
	numero++            // Incrementa 1
	numero--            // Decrementa 1
	fmt.Println(numero) // 10
	numero += 6         // Adiciona 5
	numero -= 5         // Subtrai 5
	fmt.Println(numero) // 11
	numero *= 6         // Multiplica por 5
	numero /= 5         // Divide por 5
	fmt.Println(numero) // 13
	fmt.Println("------------------------------------")

	// Operador Ternário (não existe em Go)
	// resultado := (condicao) ? valor1 : valor2
}
