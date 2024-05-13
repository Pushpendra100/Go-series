package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now() 
	fmt.Println(presentTime) // 2024-04-26 23:37:39.877434183 +0530 IST m=+0.000133675

	// we have to provide this value only to get the required format
	fmt.Println(presentTime.Format("01-02-2006")) //04-26-2024
	fmt.Println(presentTime.Format("01-02-2006 Monday")) //04-26-2024 Friday
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //04-26-2024 23:40:54 Friday

	createdDate := time.Date(2020, time.April, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate) //04-26-2024 23:40:54 Friday
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday")) //04-10-2020 23:23:00 Friday

}