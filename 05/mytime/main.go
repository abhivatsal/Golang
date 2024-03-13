package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println("The present time is ", presentTime)

	fmt.Println("formatted time is", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	fmt.Println("The created date is ", createDate)
}
