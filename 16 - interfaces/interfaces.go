package main

import (
	"fmt"
	"math"
)

type FormaGeometrica interface {
	area() float64
}

func escreverArea(f FormaGeometrica) {
	fmt.Printf("A área é: %0.2f \n", f.area())
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) area() float64 {
	return r.Largura * r.Altura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

func main() {
	// ***
	// Interfaces
	// ***
	// As interfaces são um tipo especial em Go que define um conjunto
	// de métodos que um tipo deve implementar. Elas permitem que
	// diferentes tipos sejam tratados de forma uniforme, desde que
	// implementem os métodos definidos na interface. Isso é útil para
	// abstração e polimorfismo, permitindo que diferentes tipos
	// sejam usados de forma intercambiável, desde que implementem a
	// interface esperada.

	retangulo := Retangulo{10, 15}
	escreverArea(retangulo)

	circulo := Circulo{5}
	escreverArea(circulo)
}
