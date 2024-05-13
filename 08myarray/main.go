package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	// in array you have to provide the size of array
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList) //[Apple Tomato  Peach]
	fmt.Println("Fruit list is: ", len(fruitList)) // 4

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", vegList) //[potato beans mushroom]
	fmt.Println("Vegy list is: ", len(vegList)) // 3

}