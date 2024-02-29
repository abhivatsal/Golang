package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMessage := "Hello, World!"

	println(welcomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	//comma ok || error ok

	input, _ := reader.ReadString('\n')

	fmt.Println("Hello, ", input)
	fmt.Printf("Type is %T\n", input)
}
