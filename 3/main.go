package main

import (
	"fmt"
)

// write a program that checks if given number if even of odd

func main(){
	var input int; 

	fmt.Println("Enter a number")
	fmt.Scan(&input)

	if input%2 == 0 {
		fmt.Println("Even number")
		return
	}
	fmt.Println("odd number")

}