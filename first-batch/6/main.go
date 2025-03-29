package main

import (
	"fmt"
)

// write a program that takes a year and checks if its a leap year (ggl wiki leap year for rules)

func main(){
	var year int
	fmt.Println("Enter a year to check if it is a leap year")
	fmt.Scan(&year)
	
	// I guess the secound case could be split up in two different cases, but I'll keep it as is
	switch{
		case year%100 == 0 && year%400 != 0:
			fmt.Println(year, "is NOT a leap year")
		case year%400 == 0 || year%4 == 0:
			fmt.Println(year, "is a leap year")
		default:
			fmt.Println(year, "is NOT a leap year")	
	}
}