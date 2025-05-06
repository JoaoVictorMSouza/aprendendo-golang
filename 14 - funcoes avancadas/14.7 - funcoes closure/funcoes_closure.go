package main

import "fmt"

func funcaoClosure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	// ***
	// Funções closure
	// ***
	// Funções closure são funções que capturam o ambiente
	// em que foram criadas. Isso significa que elas podem
	// acessar variáveis que estão fora do seu escopo, mas que
	// estão dentro do escopo em que a função foi criada.
	fmt.Println("Funcões closure")

	texto := "Dentro da função main"
	fmt.Println(texto)
	fmt.Println("------------------------------------")

	funcaoNova := funcaoClosure()

	funcaoNova()
	fmt.Println("------------------------------------")
}
