package main

import "fmt"

func main() {
	var firstName string = "John"

	fmt.Println(firstName)
	fmt.Printf("Variable type is : %T\n", firstName)

	var isLearning bool = true
	fmt.Println(isLearning)
	fmt.Printf("Variable type is : %T\n", isLearning)

	//default values and aliases
	var age int
	fmt.Println(age)
	fmt.Printf("Variable type is : %T\n", age)

	//implicit type
	var lastName = "Doe"
	fmt.Println(lastName)
	fmt.Printf("Variable type is : %T\n", lastName)

	// without var
	noOfUsers := 30 // This is not allowed outside a function
	fmt.Println(noOfUsers)
	fmt.Printf("Variable type is : %T\n", noOfUsers)

}
