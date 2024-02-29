package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	//comma ok || error ok
	fmt.Print("Enter your name: ")
	input, _ := reader.ReadString('\n')

	fmt.Println("Hello, ", input)

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add 1 to Rating is ", numrating+1)
	}

}
