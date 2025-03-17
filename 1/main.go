package main

import (
	"fmt"
)

// write a program that checks if number is positive or negative or 0 and prints the result ("negative", "positive", "zero")
func main(){
	var number float32
	
	fmt.Println("Please enter a number")
	fmt.Scan(&number)

	if number >= 0 {
		if number == 0 { 
			fmt.Printf("Number is zero")
		} else {
			fmt.Printf("Number is positive")
		} 
	} else { 
		fmt.Printf("Number is negative")
	}

}