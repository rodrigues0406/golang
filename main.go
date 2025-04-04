package main




 import "fmt" 

 func main() {

	var ages = [4]int{17, 16, 20, 40} 
	names := [4] string {"Julia", "Malu", "Will", "Vivi"}
	fmt.Println(ages)
	fmt.Println(names)
	names[3] = "Jujuu"
	fmt.Println(names)
	//Slice
	var score = []int{10,20,30,40}
	fmt.Println(score)
	score[1] = 2 
	fmt.Println(score, len(score), cap(score))
	rangeOne := score[1:3]
	fmt.Println(rangeOne)

	var superherois = [] string {"homem-aranha", "homem de ferro", "hulk"}
	fmt.Println(superherois)
	superherois = append(superherois, "laterna verde", "flash")
	fmt.Println(superherois, len(superherois), cap(superherois))
 }
