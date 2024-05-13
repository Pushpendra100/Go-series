package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// No inheritance in golang; No super or parent

	pushpendra := User{"Pushpendra", "pushpa@gmail.com", true, 20}
	fmt.Println(pushpendra) //{Pushpendra pushpa@gmail.com true 20}
	pushpendra.GetStatus()
	pushpendra.NewMail()
	fmt.Println(pushpendra) //{Pushpendra pushpa@gmail.com true 20}
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// methods
func (u User) GetStatus(){  
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail(){  // user is passed as a copy thus the actual mail is not changed, thats why pointers come into role
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}