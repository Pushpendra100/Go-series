package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages) //map[JS:javascript PY:Python RB:Ruby]
	fmt.Println("JS shorts for: ", languages["JS"]) //javascript
	
	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages) //map[JS:javascript PY:Python]


	//Loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

// 	For key JS, value is javascript
// 	For key PY, value is Python

}