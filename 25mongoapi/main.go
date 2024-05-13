package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pushpendra100/mongoapi/router"
)

// keep only one go file in the root directory of the project
func main(){
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}