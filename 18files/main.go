package main

import (
	"fmt"
	"io"
	"os"
)

func main (){
	fmt.Println("Welcome to files in golang")
	content := "This need to go in a file - Pushpendra Pal"

	file, err := os.Create("./myfile.txt")
	// if err != nil {
	// 	panic(err) // panic will shutdown the execution of program
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is: ", length)

	defer file.Close()

	readFile("./myfile.txt")
}

func readFile(filename string){
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	// fmt.Println("Text data inside the file is \n", databyte)   //[84 104 105 115 32 110 101 101 100 32 116 111 32 103 111 32 105 110 32 97 32 102 105 108 101 32 45 32 80 117 115 104 112 101 110 100 114 97 32 80 97 108]
	fmt.Println("Text data inside the file is \n", string(databyte)) //This need to go in a file - Pushpendra Pal
}

func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}