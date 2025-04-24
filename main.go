package main 

import "fmt"  

func analisarNotas (nota1, nota2 float64) (float64, string){
media := (nota1 + nota2) / 2 

var texto string
if media >= 6 {
texto = "Aprovado"
} else { 
texto = "Reprovado"
}

return media, texto 

}


func main(){
media, resultado := analisarNotas(7.5, 5.5)
fmt.Println("media", media)
fmt.Println("resultado", resultado)
}

