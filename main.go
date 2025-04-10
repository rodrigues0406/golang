package main

import (
	"fmt"
)

func main() {
age := 45
fmt.Println(age <= 50)
fmt.Println(age >= 50)
fmt.Println(age == 50)
fmt.Println(age != 50)

if age < 50 {
	fmt.Println("menor que 30 anos ")
} else if age < 40 {
	fmt.Println(" menor que 40 anos ")
} else {
  fmt.Println("não é menor que 40 anos")
}

names := []string{"julia", "maria", "silva", "rodrigues"}

for index, value := range names {
	if index == 1 {
	fmt.Println("Continue após a posição", index, "é valor", value)
	continue 

}
if index > 2 {
	fmt.Println("sair apos ", index)
	break
}

	fmt.Println("valor: ", value)


}
}

 
