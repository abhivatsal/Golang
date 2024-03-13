package main

import "fmt"

func main() {
	fmt.Println("Maps in go lang")
	languages := make(map[string]string)

	languages["JA"] = "Java"
	languages["RB"] = "Ruby"
	languages["Js"] = "Javascript"

	fmt.Println(languages)

	//delete the values of a specific key

	delete(languages, "JA")

	fmt.Println("After delete the key", languages)

	// loops for map

	for key, value := range languages {
		fmt.Printf("For key %v value is %v\n", key, value)
	}
}
