package main

import (
	"errors"
	"fmt"
)

func main() {
	// ***
	// Numeros inteiros
	// ***
	var (
		numeroInt int = 1000 // Pela arquitetura do computador, o tipo int pode ser 32 ou 64 bits. Inferência de tipo
		// int8 = byte
		numeroInt8  int8  = 100  // 8 bits; Também pode ser byte, que é um alias para int8
		numeroInt16 int16 = 1000 // 16 bits
		// int32 = rune
		numeroInt32 int32 = 10000  // 32 bits; Também pode ser rune, que é um alias para int32
		numeroInt64 int64 = 100000 // 64 bits
	)
	fmt.Println("Numeros inteiros")
	fmt.Println(numeroInt, numeroInt8, numeroInt16, numeroInt32, numeroInt64)

	// ***
	// Numeros naturais
	// ***
	var (
		numeroUint   uint   = 1000   // Pela arquitetura do computador, o tipo uint pode ser 32 ou 64 bits. Inferência de tipo
		numeroUint8  uint8  = 100    // 8 bits
		numeroUint16 uint16 = 1000   // 16 bits
		numeroUint32 uint32 = 10000  // 32 bits
		numeroUint64 uint64 = 100000 // 64 bits
	)
	fmt.Println("Numeros naturais")
	fmt.Println(numeroUint, numeroUint8, numeroUint16, numeroUint32, numeroUint64)

	// ***
	// Numeros reais
	// ***
	var (
		numeroFloat32 float32 = 100.0  // 32 bits
		numeroFloat64 float64 = 1000.0 // 64 bits
	)
	fmt.Println("Numeros reais")
	fmt.Println(numeroFloat32, numeroFloat64)

	// ***
	// Textos
	// ***
	var texto string = "Texto"
	fmt.Println(texto)

	// ***
	// Caracteres
	// ***
	char := 'A' // 32 bits; O tipo char é um alias para int32
	fmt.Println(char)

	// ***
	// Booleano
	// ***
	var booleanoT bool = true // 1 byte
	booleanoF := false        // 1 byte
	fmt.Println(booleanoT, booleanoF)

	// ***
	// Erro
	// ***
	var erro error
	fmt.Println(erro) // nil
	erro = errors.New("Erro interno")
	fmt.Println(erro) // Erro interno
}
