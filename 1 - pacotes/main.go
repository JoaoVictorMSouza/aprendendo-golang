package main

import (
	"fmt"
	"modulo_exemplo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("joao.victor@gmail.com")
	fmt.Println(erro)

	erro2 := checkmail.ValidateFormat("joaovictor.com")
	fmt.Println(erro2)
}
