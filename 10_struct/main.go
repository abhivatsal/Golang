package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")
	user := User{"Abhishek Kumar", "25", "ac08@yopmail.com", "active"}

	fmt.Println(user)

	fmt.Printf("user data is %+v\n", user) // It will give the struct of user with their keys
}

type User struct {
	Name   string
	Age    string
	Email  string
	Status string
}
