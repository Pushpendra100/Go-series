package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// No inheritance in golang; No super or parent

	pushpendra := User{"Pushpendra", "pushpa@gmail.com", true, 20}
	fmt.Println(pushpendra) //{Pushpendra pushpa@gmail.com true 20}
	fmt.Printf("Pushpendra details are: %+v\n", pushpendra) //{Name:Pushpendra Email:pushpa@gmail.com Status:true Age:20}
	fmt.Printf("Name is %v and email is %v\n", pushpendra.Name, pushpendra.Email) //Name is Pushpendra and email is pushpa@gmail.com
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
