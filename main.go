package main 

import "fmt" 

func main(){ 
    estoque := map[string]int{
        "coxinha" : 10,
        "pão de queijo" : 15,
        "refrigerante" : 20,
        
    }
  
 estoque["coxinha"] -=2

 estoque["pão de queijo"]-=1

 estoque["refrigerante"] -=3

 for k,v := range estoque{
    fmt.Println(" produto/ estoque", k,v)
 }
}






