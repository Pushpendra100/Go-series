package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr) //nil

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is ", ptr) //0xc000112010
	fmt.Println("Value of actual pointer is ", *ptr) //23

	*ptr = *ptr * 2
	fmt.Println("New value is: ", myNumber) //46

}

// pointers give you the gurantee that the operation will be performed on the actual value at that address