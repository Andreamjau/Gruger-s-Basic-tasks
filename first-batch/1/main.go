package main

import (
	"fmt"
)

// write a program that checks if number is positive or negative or 0 and prints the result ("negative", "positive", "zero")
func main(){
	var number float32
	
	fmt.Println("Please enter a number")
	fmt.Scan(&number)

	switch {
		case number > 0:
			fmt.Printf("Number is positive")
		case number < 0:
			fmt.Printf("Number is negative")
		default:
			fmt.Printf("Number is zero")
	}
}