package main

import "fmt"

func main() {

	fmt.Println("Welcome to slices study of golang")

	var fruitlist = []string{"Apple", "Mango"}

	fmt.Printf("Fruit list is a type %T", fruitlist)

	//to add the data in slices

	fruitlist = append(fruitlist, "Banana")

	fmt.Println("Fruit list is ", fruitlist)
}
