package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=djalf17483"

func main(){
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	//parsing
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme) //https
	fmt.Println(result.Host) //lco.dev:3000
	fmt.Println(result.Path) // /learn
	fmt.Println(result.Port()) //3000
	fmt.Println(result.RawQuery) //coursename=reactjs&paymentid=djalf17483

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams) // url.Values
	fmt.Println(qparams) //map[coursename:[reactjs] paymentid:[djalf17483]]
	fmt.Println(qparams["coursename"]) //[reactjs]

	for t, val := range qparams {
		fmt.Println("Param is:  ",t , val)
	}


	// we always pass reference here
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=pushpendra",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL) //https://lco.dev/tutcss
}