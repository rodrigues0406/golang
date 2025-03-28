package main

import "fmt"

func main(){

var numero1 int
var numero2 int 
fmt.Println("digite um numero:")
fmt.Scan(&numero1)
fmt.Println("digite um numero:")
fmt.Scan(&numero2)

fmt.Println("A soma é:", numero1 + numero2)
fmt.Println(" A subtração é:", numero1 - numero2)
fmt.Println("A multiplicação é:", numero1 * numero2)
fmt.Println("A divisão é:", numero1 / numero2)
fmt.Println("O resto da divisão é:", numero1 % numero2)

}
