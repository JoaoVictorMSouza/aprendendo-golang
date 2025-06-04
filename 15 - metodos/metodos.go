package main

import "fmt"

type Usuario struct {
	Nome  string
	Idade uint8
}

func (u Usuario) salvar() {
	fmt.Println("Salvando usuário:", u.Nome, "com idade", u.Idade)
}

func (u Usuario) maiorDeIdade() bool {
	return u.Idade >= 18
}

func (u *Usuario) atualizarIdade(novaIdade uint8) {
	u.Idade = novaIdade
	fmt.Println("Idade atualizada para:", u.Idade)
}

func main() {
	// ***
	// Métodos
	// ***
	fmt.Println("Métodos")

	// Métodos são funções que estão associadas a alguma estrutura
	// ou tipo. Eles são definidos com um receptor, que é o tipo
	// ou estrutura que o método está associado. O receptor é definido
	// entre parênteses antes do nome do método. Isso permite que o
	// método acesse os campos e métodos do receptor, e também
	// que o método seja chamado em uma instância do receptor.

	var usuario Usuario = Usuario{
		Nome:  "João",
		Idade: 30,
	}

	usuario.salvar()
	fmt.Println("É maior de idade?", usuario.maiorDeIdade())

	usuario.atualizarIdade(usuario.Idade + 1)
}
