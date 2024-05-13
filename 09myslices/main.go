package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	// we can add as many values we want
	var fruitList = []string{"Apple", "tomato", "peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList) //[]string

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList) //[Apple tomato peach Mango Banana]
	
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList) //[tomato peach Mango Banana]

	fruitList = append(fruitList[1:2])
	fmt.Println(fruitList) //[peach]





	// highScore := make([]int, 4)
	// highScore[0] = 123
	// highScore[1] = 52
	// highScore[2] = 523
	// highScore[3] = 230

	// fmt.Println(highScore) //[123 52 523 230]




	highScore := make([]int, 4)
	highScore[0] = 123
	highScore[1] = 52
	highScore[2] = 523
	highScore[3] = 230
	// highScore[4] = 230  // -------> runtime error

	highScore = append(highScore, 230, 666) //[123 52 523 230 230 666]

	fmt.Println(highScore) 


	sort.Ints(highScore)
	fmt.Println(highScore) //[52 123 230 230 523 666]
	fmt.Println(sort.IntsAreSorted(highScore)) //true





	// How to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses) //[reactjs javascript swift python ruby]
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) //[reactjs javascript python ruby]

}