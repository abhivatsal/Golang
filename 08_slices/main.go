package main

import "fmt"

func main() {

	fmt.Println("Welcome to slices study of golang")

	var fruitlist = []string{"Apple", "Mango"}

	fmt.Printf("Fruit list is a type %T", fruitlist)

	//to add the data in slices

	fruitlist = append(fruitlist, "Banana")

	fmt.Println("Fruit list is ", fruitlist)

	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 234
	highscores[2] = 234
	highscores[3] = 234
	// highscores[4] = 234

	highscores = append(highscores, 12, 10, 12)
	fmt.Println(highscores)

	// Remove the element from slices based on Index
	var courses = []string{"react", "data", "js", "vue js"}
	var index = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
