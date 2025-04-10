package main

import (
	"fmt"
)

func main() {
	var saldo int
	var comando string
	var valor int

	fmt.Println("Digite seu saldo inicial: ")
	fmt.Scanln(&saldo)

	fmt.Println("saldo atual:", saldo)
	fmt.Println(" deseja sacar ou depositar? sacar/depositar")
	fmt.Scanln(&comando)
	if comando == "sacar" {
		fmt.Println("digite o valor que irá sacar")
		fmt.Scanln(&valor)
		if valor > saldo {
			fmt.Println("saldo indisponivel", saldo)
		} else {
			saldo -= valor
			fmt.Println("saque realizado com sucesso!", valor, "seu saldo atual agora é", saldo)
		}
	} else if comando == "depositar" {
		fmt.Println("digite o valor que irá depositar:")
		fmt.Scan(&valor)
		saldo += valor
		fmt.Println("deposito realizado com sucesso, seu saldo é", saldo)

	}

}
