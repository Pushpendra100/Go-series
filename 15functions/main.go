package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, _ := proAdder(2, 5, 10)
	fmt.Println("pro result is: ", proRes)

}

func proAdder(values ...int) (int, string) { //profunctions
	total := 0

	for _, val := range values{
		total += val
	}
	return total, "Hi pro result function"
}

func adder(valOne int, valTwo int) int { // function signature - what values type is taken in input and what type is given in return
	return valOne + valTwo
}

func greeterTwo(){
	fmt.Println("Another method")
}

func greeter(){
	fmt.Println("Namastey from golang")
}