package main

import "fmt"

// Structs são tipos de dados compostos que permitem
// agrupar diferentes tipos de dados em uma única estrutura.
// Eles são úteis para representar entidades complexas
// com várias propriedades.
type Usuario struct {
	Nome     string
	Idade    uint8
	Endereco Endereco
}

type Endereco struct {
	Logradouro string
	Numero     uint8
}

func main() {
	fmt.Println("Structs")

	var endereco1 Endereco = Endereco{
		Logradouro: "Rua A",
		Numero:     123,
	}
	fmt.Println(endereco1)
	fmt.Println("------------------------------------")

	endereco2 := Endereco{
		Logradouro: "Rua B",
		Numero:     200,
	}
	fmt.Println(endereco2)
	fmt.Println("------------------------------------")

	var u Usuario
	u.Nome = "Lucas"
	u.Idade = 25
	u.Endereco = endereco1
	fmt.Println(u)
	fmt.Println("------------------------------------")

	u2 := Usuario{
		Nome:  "Lucas",
		Idade: 25,
		Endereco: Endereco{
			Logradouro: "Rua A",
			Numero:     123,
		},
	}
	fmt.Println(u2)
	fmt.Println("------------------------------------")

	u3 := Usuario{"Lucas", 25, Endereco{"Rua A", 123}}
	fmt.Println(u3)
	fmt.Println("------------------------------------")

	u4 := Usuario{"João", 30, endereco2}
	fmt.Println(u4)
	fmt.Println("------------------------------------")

	u5 := Usuario{
		Nome:     "Maria",
		Idade:    28,
		Endereco: endereco1,
	}
	fmt.Println(u5)
	fmt.Println("------------------------------------")
}
