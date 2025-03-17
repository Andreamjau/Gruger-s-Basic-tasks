package main

import (
	"fmt"
)

// write a program that categorizes given number (0-10 -> small, 11-100 -> medium, 101-1000 -> large)

func main(){
	var input int

	fmt.Println("Enter a number")
	fmt.Scan(&input)

	if input > 10 {
		if input > 100 {
			fmt.Println("Number is large")
			return
		}
		fmt.Println("Number is medium")
			return
	}

	fmt.Println("Nuber is small")
}