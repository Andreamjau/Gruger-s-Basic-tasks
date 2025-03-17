package main

import (
	"fmt"
)

// write a program that checks if given number if even of odd

func main(){
	var input float32; 

	fmt.Println("Enter a number")
	fmt.Scan(&input)

	// Checking ifthe converted float to interger is the same as the input
	// Had to google for the float int converting
	floatToInt := float32(int32(input/2))
	if input/2 == floatToInt{
		fmt.Println("Even number")
		return
	}
	fmt.Println("odd number")
	
}