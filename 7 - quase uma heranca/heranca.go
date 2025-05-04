package main

import "fmt"

type Pessoa struct {
	Nome      string
	SobreNome string
	Idade     uint8
	Altura    uint8
}

type Estudante struct {
	Pessoa
	Curso     string
	Faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := Pessoa{
		Nome:      "Lucas",
		SobreNome: "Silva",
		Idade:     25,
		Altura:    180,
	}
	fmt.Println(p1)
	fmt.Println("------------------------------------")

	e1 := Estudante{
		Pessoa:    p1,
		Curso:     "Engenharia",
		Faculdade: "UFMG",
	}
	fmt.Println(e1)
	fmt.Println(e1.Nome)
	fmt.Println(e1.Pessoa.Nome)
	fmt.Println("------------------------------------")
}
