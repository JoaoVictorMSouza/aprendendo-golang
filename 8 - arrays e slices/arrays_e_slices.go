package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	// ***
	// Arrays
	// ***
	fmt.Println("Arrays")

	var array1 [5]int
	fmt.Println(array1)
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println(array1)
	fmt.Println("------------------------------------")

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)
	fmt.Println("------------------------------------")

	array3 := [...]string{"Lucas", "João", "Maria"}
	fmt.Println(array3)
	fmt.Println("------------------------------------")

	// ***
	// Slices
	// ***
	fmt.Println("Slices")

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	fmt.Println("------------------------------------")
	slice1 = append(slice1, 6)
	fmt.Println(slice1)
	fmt.Println("------------------------------------")

	// Criando um slice a partir de um array
	slice2 := array2[2:4]
	fmt.Println(slice2)
	fmt.Println("------------------------------------")
	array2[2] = 33
	fmt.Println(slice2)
}
