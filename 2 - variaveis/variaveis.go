package main

import "fmt"

func main() {
	// Declarando uma variavel explícita do tipo string
	var variavel1 string = "Variavel 1"
	fmt.Println(variavel1)

	// Declarando uma variavel implícita do tipo string
	// Conhecido com inferência de tipo
	variavel2 := "Variavel 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	// Invertendo os valores das variaveis sem usar uma variavel temporaria
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

	// Constantes
	const constante1 string = "Constante 1"
	fmt.Println(constante1)
}
