package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// ***
	// Maps
	// Maps são estruturas de dados que armazenam pares de chave-valor.
	// Eles são semelhantes a dicionários em outras linguagens de programação.
	// ***
	usuario := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])
	fmt.Println("------------------------------------")

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Lucas",
			"ultimo":   "Silva",
		},
		"curso": {
			"nome":   "Programação em Go",
			"campus": "São Paulo",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["primeiro"])
	fmt.Println("------------------------------------")

	delete(usuario2, "nome")
	fmt.Println(usuario2)
	fmt.Println("------------------------------------")

	usuario2["signo"] = map[string]string{
		"nome": "Áries",
	}
	fmt.Println(usuario2)

	fmt.Println("------------------------------------")
}
