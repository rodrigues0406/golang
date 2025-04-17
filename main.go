package main

import (
	"fmt"
)
func main() {
	alunoIdade :=  make(map[string]int)
    alunoIdade["Julia"] = 16
    alunoIdade["Isadora"] = 16
    alunoIdade["Isabelly"] = 16
    alunoIdade["Manu"] = 15
    fmt.Println("Idade da Julia", alunoIdade["Julia"])
    notasAlunos := map[string]float64{
        "Bruno" : 9.7,
        "Otávio" : 10,
        "Fabiano" : 8.7,
        "Isabelly" :9.5,
    
    }
    for nome,nota := range notasAlunos{
        fmt.Printf("%s tirou a nota %.1f\n", nome,nota)
    }
}



