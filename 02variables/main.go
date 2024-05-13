package main

import "fmt"

const LoginToken string = "ghabbhhjd" // Public - captial letter means that this variable is accessible from any other file

func main() {
	var username string = "pushpendra"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	// var smallVal uint8 = 256 //-gives error
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	// var smallFloat float32 = 255.4555471203  //-gives 255.45555
	var smallFloat float64 = 255.4555471203 //-gives 255.4555471203
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable) //-gives 0  - doesn't put garbage value in it
	fmt.Sprintf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "pushpendra.com"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000.0 // this syntax is not allowed outside the function
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
