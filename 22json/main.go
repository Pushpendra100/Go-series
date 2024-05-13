package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeonline.in", "bcd123", []string{"full-stack", "js"}},
		{"Web3 Bootcamp", 499, "LearnCodeonline.in", "edf123", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}


func DecodeJson(){
	jsonDataFromWeb := []byte(`
        {
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "Platform": "LearnCodeonline.in",
                "tags": ["web-dev", "js"]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%v\n", lcoCourse) //{ReactJS Bootcamp 299 LearnCodeonline.in  [web-dev js]}
		fmt.Printf("%#v\n", lcoCourse) //main.course{Name:"ReactJS Bootcamp", Price:299, Platform:"LearnCodeonline.in", Password:"", Tags:[]string{"web-dev", "js"}}
	} else {
		fmt.Println("JSON was not valid")
	}


	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	// map[string]interface{} -- value type is not fixed thus interface is used
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData) //map[string]interface {}{"Platform":"LearnCodeonline.in", "Price":299, "coursename":"ReactJS Bootcamp", "tags":[]interface {}{"web-dev", "js"}}

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is: %T\n", k, v, v)
	}
	/*
	Key is coursename and value is ReactJS Bootcamp and type is: string
	Key is Price and value is 299 and type is: float64
	Key is Platform and value is LearnCodeonline.in and type is: string
	Key is tags and value is [web-dev js] and type is: []interface {}
	*/

}