package main

import "fmt"

func main() {

	fmt.Println("Welcome to array study of golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Banana"

	fmt.Println("Fruit list is ", fruitList)
	fmt.Println("Fruit list length", len(fruitList))

}
