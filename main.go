package main

import "fmt"

func main(){

 var usuario string
 var senha string 
 fmt.Println("Digite o seu nome:")
 fmt.Scanln(&usuario)

 fmt.Println("Digite o numero")
 fmt.Scanln(&senha)

if usuario == "julia" && senha == "1234" {
	fmt.Println("Login feito com sucesso ")
	
}else{
fmt.Println("Acesso negado")
}

}
