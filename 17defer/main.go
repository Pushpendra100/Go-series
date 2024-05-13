package main

import "fmt"

func main(){
	defer fmt.Println("end")
	defer fmt.Println("world")
	fmt.Println("Hello")

	myDefer()
}

// execution happens line by line and the statements with defer keyword will go to the end of the function in LIFO order

/*
Hello
world
end
*/


func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

/*
Hello
4 3 2 1 0world
end
*/