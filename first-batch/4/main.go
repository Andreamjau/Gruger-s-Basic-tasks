package main

import (
	"fmt"
)

// write a program that categorizes given number (0-10 -> small, 11-100 -> medium, 101-1000 -> large)

func main(){
	var input int

	fmt.Println("Enter a number")
	fmt.Scan(&input)
 
	switch {
	case input <= 10:
		fmt.Println("Number is small")
	case input > 10 && input <= 100:
		fmt.Println("Number is medium")
	case input > 100: 
		fmt.Println("Number is large")
	default: 
		fmt.Println("waht?")
	}
}