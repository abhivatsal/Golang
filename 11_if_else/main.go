package main

import "fmt"

func main() {
	fmt.Println("If else statement exercise")

	count := 10

	var result string

	if count%2 == 0 {
		result = "Even Number"
	} else if count%2 != 0 {
		result = "Odd number"
	}

	if num := 11; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}

	fmt.Printf("The number %v is %v", count, result)
}
