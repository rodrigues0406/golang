package main

import (
	"fmt"
)

func main() {
	var numeros [5]int
	var soma int

	//slice
	for i := 0; i < 5; i++ {
		fmt.Printf("Digite o  numero %d: ", i+1)
		fmt.Scanln(&numeros[i])
		soma += numeros[i]
	}

	//slice
	fmt.Printf("A soma é: %d\n", soma)
}

 
